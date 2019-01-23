package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
)

/*
[{
  "weather": "sunny",
  "image_name": "rec_20170517_135411_00004684.jpg",
  "image_uuid": "30c59b90-be92-4f5b-932b-a8e45a9953b8",
  "package_image_uuid": "000057e5-5b92-44e0-a8aa-915ee1223df0"
  "scene_category": "highway",
  "annotations": [
    {
      "anno": [
        {
          "data": {
            "point2d": [
              [
                515.2702557990351,
                323.9477503628448
              ],
              [
                844.2103829285265,
                620.4442249378596
              ]
            ]
          },
          "index": "620ee668-8d82-4635-b5aa-e903c5ff413e",
          "label": [
            "huakuang",
            "a"
          ],
          "type": "rect"
        }
      ],
      "anno_uuid": "c277ebee-1c8f-4fa6-9cfb-0c30e1465b26",
      "userid": 858
    },
    {
      "anno": [
        {
          "data": {
            "point2d": [
              [
                7.505667733763858,
                52.05849478390462
              ],
              [
                211.23093479306857,
                172.23733233979135
              ]
            ]
          },
          "index": "044cf7ff-f3be-4310-9a3e-ac32d61bfc62",
          "label": [
            "huakuang",
            "a"
          ],
          "type": "rect"
        }
      ],
      "anno_uuid": "2fa6cd67-a502-4a36-9ca8-a50170ce3e7e",
      "userid": 912
    }
  ],
}]
*/

// TResultJSONOneImageV2 Result.json 第二版协议
type TResultJSONOneImageV2 struct {
	ImageName        string                  `json:"image_name"`
	ImageUUID        string                  `json:"image_uuid"`
	PackageImageUUID string                  `json:"package_image_uuid"`
	Annotations      []TResultJSONAnnotation `json:"annotations"`
}

type TResultJSONAnnotation struct {
	Anno     []interface{} `json:"anno"`
	AnnoUUID string        `json:"anno_uuid"`
	UserID   int           `json:"userid"`
}

// FileIsV2Version 判断 result.json 是否符合该版本
func FileIsV2Version(filename string) bool {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		beego.Info(err)
		return false
	}

	return DataIsV2Version(data)
}

// DataIsV2Version 判断 result.json 是否符合该版本
func DataIsV2Version(data []byte) bool {
	obj := make([]TResultJSONOneImageV2, 0)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		beego.Info(err)
		return false
	}

	if len(obj) == 0 {
		return false
	}

	oneImage := obj[0]

	if len(oneImage.ImageUUID) > 0 {
		return true
	}

	if len(oneImage.PackageImageUUID) > 0 {
		return true
	}

	if len(oneImage.Annotations) == 0 {
		return false
	}

	oneAnno := oneImage.Annotations[0]

	if len(oneAnno.AnnoUUID) > 0 {
		return true
	}

	if oneAnno.UserID > 0 {
		return true
	}

	if len(oneAnno.Anno) > 0 {
		return true
	}

	return false
}

func test(n int) int {
	i, j := 1, 2
	for idx := 2; idx < n; idx++ {
		i, j = j, i+j
	}
	return j
}

func main() {
	fmt.Println("ret = ", test(4))
}
