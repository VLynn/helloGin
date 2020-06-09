package main

import (
    "strconv"
    "log"
    "time"
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "helloGin/model"
)

func setupRouter() *gin.Engine {
    router := gin.Default()
    router.Static("/front", "./front")

    pg := router.Group("/user/profile")
    {
        pg.GET("/get_list", func(c *gin.Context) {
            offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
            num, _ := strconv.Atoi(c.DefaultQuery("num", "10"))

            users := model.GetList(offset, num)
            c.JSON(http.StatusOK, gin.H{
                "list": users,
                "count": len(users),
            })
        })

        pg.POST("/insert", func(c *gin.Context) {
            var user model.UserProfile
            err := c.BindJSON(&user)
            if err != nil {
                log.Println(err)
                return
            }
            log.Println(user.Name)

            insert_id := model.Insert(user)
            c.JSON(http.StatusOK, gin.H{
                "insert_id": insert_id,
            })
        })

        pg.POST("/update", func(c *gin.Context) {
            data, err := c.GetRawData()
            if err != nil {
                c.String(http.StatusNotAcceptable, "get post data failed, err = %s", err)
                return
            }

            var user map[string]interface{}
            json.Unmarshal(data, &user)
            id := user["id"].(float64)
            name := user["name"].(string)
            company := user["company"].(string)

            log.Println(company)
            model.Update(int(id), name, company)
            c.String(http.StatusOK, "success.")
        })

        pg.GET("/delete", func(c *gin.Context) {
            id_str := c.Query("id")
            if id_str == "" {
                c.String(http.StatusNotAcceptable, "param id missing")
                return
            }

            id, _ := strconv.Atoi(id_str)
            model.Delete(id)
            c.String(http.StatusOK, "deleted")
        })
    }
 
    router.Any("/", func(c *gin.Context) {
        c.String(http.StatusOK, "hello world")
    })

    // URL query param
    router.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "none")
        lastname := c.Query("lastname")
        c.JSON(http.StatusOK, gin.H{"firstname": firstname, "lastname": lastname})

        // 绑定URL查询参数，并添加参数验证
        var person Person
        if err := c.ShouldBindQuery(&person); err != nil {
            log.Println("bind query failed, err =", err)
        }
        log.Printf("person.firstname = %s\n", person.Name)
        log.Printf("person.lastname = %s\n", person.Lastname)
    })

    // POST JSON 字符串
    router.POST("/set", func(c *gin.Context) {
        data, err := c.GetRawData()
        if err != nil {
            c.String(http.StatusNotAcceptable, "get post data failed, err = %s", err)
            return
        }
        log.Println(string(data))
        c.String(http.StatusOK, "OK")
    })

    // 上传文件
    router.POST("/upload", func(c *gin.Context) {
        fh, err := c.FormFile("file")
        if err != nil {
            c.String(http.StatusNotAcceptable, "get file failed, err = %s", err)
            return
        }
        c.SaveUploadedFile(fh, "front/" + fh.Filename)
        c.String(http.StatusOK, "filename = %s, size = %d", fh.Filename, fh.Size)
    })

    // 异步处理
    router.GET("/async", func(c *gin.Context) {
        c.String(http.StatusOK, "start handle...")
        // 在goroutine中使用context时，必须使用只读副本
        cCopy := c.Copy()

        go func() {
            time.Sleep(5 * time.Second)
            log.Println("done! url = ", cCopy.Request.URL)
        }()
    })

    return router
}

func main() {
    router := setupRouter()
    router.Run()
}