// ## login
// 	open "http://haue.dr2am.cn/home"
// 	click "body > div.position-relative.js-header-wrapper > header > div > div.HeaderMenu.HeaderMenu--bright.d-lg-flex.flex-justify-between.flex-auto > div > span > div > a"
// 	type "#inputUserName" "hauelib"
// 	type "#inputPassword" "test2019"
// 	click "form > button"
## search
    // open "http://www.wanfangdata.com.cn.haue.dr2am.cn"
    open "http://www.wanfangdata.com.cn"
    type "#keyWords" "asp"
    click "#subform > div.searchBox > div.pubsearchBtn > input[type=\"button\"]"
    find "#here > div.right_con > div.BatchOper.clear > div.BatchOper_result_show > span"
    text 
    
## detail
    findAll ".ResultList  .title > a:nth-child(2)"
    random
    click

// ## logout
// 	open "http://haue.dr2am.cn/home"
// 	click "body > div.position-relative.js-header-wrapper > header > div > div.HeaderMenu.HeaderMenu--bright.d-lg-flex.flex-justify-between.flex-auto > div > span > div > a"
