package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RecLogs struct {
	Id        int     `orm:"column(id);pk"`
	Userid    string  `orm:"column(userid);size(255);null"`
	ImgName   string  `orm:"column(img_name);size(64);null"`
	ImgBase64 string  `orm:"column(img_base64);null"`
	RecogType string  `orm:"column(recog_type);size(50);null"`
	RecogName string  `orm:"column(recog_name);size(255);null"`
	Score     float64 `orm:"column(score);null"`
}

func (t *RecLogs) TableName() string {
	return "rec_logs"
}

func init() {
	orm.RegisterModel(new(RecLogs))
}

// AddRecLogs insert a new RecLogs into database and returns
// last inserted Id on success.
func AddRecLogs(m *RecLogs) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRecLogsById retrieves RecLogs by Id. Returns error if
// Id doesn't exist
func GetRecLogsById(id int) (v *RecLogs, err error) {
	o := orm.NewOrm()
	v = &RecLogs{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRecLogs retrieves all RecLogs matches certain condition. Returns empty list if
// no records exist
func GetAllRecLogs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RecLogs))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []RecLogs
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRecLogs updates RecLogs by Id and returns error if
// the record to be updated doesn't exist
func UpdateRecLogsById(m *RecLogs) (err error) {
	o := orm.NewOrm()
	v := RecLogs{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRecLogs deletes RecLogs by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRecLogs(id int) (err error) {
	o := orm.NewOrm()
	v := RecLogs{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RecLogs{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
