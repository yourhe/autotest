## search
    open "http://www.wanfangdata.com.cn"
    type "#keyWords" "asp"
    click "#subform > div.searchBox > div.pubsearchBtn > input[type=\"button\"]"
    find "#here > div.right_con > div.BatchOper.clear > div.BatchOper_result_show > span"
    text
    
## detail
    findAll ".ResultList  .title > a:nth-child(2)"
    random
    click
    wait 10
