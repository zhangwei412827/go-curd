package blog

import (
    controller "blog/blog/controllers"
    "net/http"

    "blog/blog/config"
    "blog/blog/utils"
)

func init(){
    //映射view函数
    controller.FuncMap["timeFormat"]=utils.TimeFormat
}

func parseConfig(){

}


func Run(){
    //parse config.json

    //set static dir  
    http.Handle("/static/",http.FileServer(http.Dir(config.CFG_STATIC_DIR)))
    
    http.HandleFunc("/",controller.IndexController)
    http.HandleFunc("/new",controller.NewController)
    http.HandleFunc("/update",controller.UpdateController)
    http.HandleFunc("/delete",controller.DeleteController)
   
    http.ListenAndServe(":9000",nil)
}
