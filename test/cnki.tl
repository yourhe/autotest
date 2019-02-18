## search
    open "http://www.cnki.net"
    activeElement
    click
    wait 1
    type "#txt_SearchText" "golang"
    click "body > div.wrapper.section1 > div.searchmain > div > div.input-box > input.search-btn"
    iframe "iframeResult"
    find ".pagerTitleCell"
    text
    
## detail
    findAll ".fz14"
    random
    click
