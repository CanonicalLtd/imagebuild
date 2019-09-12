(window.webpackJsonpwebapp=window.webpackJsonpwebapp||[]).push([[0],{20:function(e,a,t){e.exports=t(51)},25:function(e,a,t){},51:function(e,a,t){"use strict";t.r(a);var n=t(0),s=t.n(n),r=t(17),l=t.n(r),c=(t(25),t(1)),i=t(2),o=t(4),m=t(3),u=t(5),d={add:"+","add-application":"Add application",applications:"Applications",board:"Board",boards:"Boards",build:"Let's go!","choose-applications":"Choose the Pre-installed Applications","choose-applications-desc":"Select the applications that will be pre-installed on the device. For Ubuntu Core this will be snaps, for Ubuntu Classic it will be snaps or packages.","choose-board":"Choose Board","choose-board-desc":"Select the board for your device.","choose-os":"Choose the Operating System","choose-os-desc":"Select the version of Ubuntu to install on the device",classic:"Classic",confirm:"Confirm","confirm-configuration":"Confirm board configuration",copyright:"\xa9 2019 Canonical Ltd. Ubuntu and Canonical are registered trademarks of Canonical Ltd.",core:"Core","find-snaps":"Find snaps to install","get-started":"Get started",home:"Home",legal:"Legal information",os:"Operating System","pre-installed-snaps":"Pre-installed snaps",privacy:"Data privacy","ready-to-build":"Ready to build the Ubuntu image?",remove:"X","remove-application":"Remove application","report-a-bug":"Report a bug on this site","select-board":"Select board","site-description1":"Choose your board","site-description2":"Select the Operating System","site-description3":"Pick your pre-installed applications","search-store":"enter snap name","snap-store":"Snap Store",subtitle:"Image building service for IoT devices.",title:"Ubuntu Image Builder"};function p(e){return d[e]||e}function h(){var e=localStorage.getItem("board");return JSON.parse(e)}var b=function(e){function a(){return Object(c.a)(this,a),Object(o.a)(this,Object(m.a)(a).apply(this,arguments))}return Object(u.a)(a,e),Object(i.a)(a,[{key:"render",value:function(){return s.a.createElement("div",{id:"footer"},s.a.createElement("div",{className:"row footer"},s.a.createElement("p",null,p("copyright")),s.a.createElement("nav",{className:"p-footer__nav",role:"navigation"},s.a.createElement("ul",{className:"p-footer__links"},s.a.createElement("li",{className:"p-footer__item"},s.a.createElement("a",{className:"p-footer__link",href:"https://ubuntu.com/legal"},p("legal"))),s.a.createElement("li",{className:"p-footer__item"},s.a.createElement("a",{className:"p-footer__link",href:"https://ubuntu.com/legal/data-privacy"},p("privacy"))),s.a.createElement("li",{className:"p-footer__item"},s.a.createElement("a",{className:"p-footer__link",href:"https://github.com/slimjim777/imagebuild/issues/new"},p("report-a-bug")))))))}}]),a}(n.Component),f=function(e){function a(){return Object(c.a)(this,a),Object(o.a)(this,Object(m.a)(a).apply(this,arguments))}return Object(u.a)(a,e),Object(i.a)(a,[{key:"render",value:function(){return s.a.createElement("header",{id:"navigation",class:"p-navigation header-slim"},s.a.createElement("div",{className:"p-navigation__banner row"},s.a.createElement("div",{className:"p-navigation__logo"},s.a.createElement("div",{className:"u-vertically-center"},s.a.createElement("img",{src:"/static/images/logo.png",width:"150px"})))))}}]),a}(n.Component),g=function(e){function a(){return Object(c.a)(this,a),Object(o.a)(this,Object(m.a)(a).apply(this,arguments))}return Object(u.a)(a,e),Object(i.a)(a,[{key:"render",value:function(){return s.a.createElement("div",null,s.a.createElement(f,null),s.a.createElement("section",{className:"p-strip--image is-dark header"},s.a.createElement("div",{className:"row"},s.a.createElement("div",{className:"col-5 title"},s.a.createElement("h1",null,p("title")),s.a.createElement("p",null,p("subtitle"))))))}}]),a}(n.Component),v=function(e){function a(e){var t;return Object(c.a)(this,a),(t=Object(o.a)(this,Object(m.a)(a).call(this,e))).state={token:e.token||{}},t}return Object(u.a)(a,e),Object(i.a)(a,[{key:"render",value:function(){return s.a.createElement("div",{className:"row"},s.a.createElement("section",{className:"row"},s.a.createElement("div",{className:"row"},s.a.createElement("div",{className:"first"},s.a.createElement("h2",null,p("get-started")),s.a.createElement("ul",{className:"p-list"},s.a.createElement("li",{className:"p-list__item is-ticked"},p("site-description1")),s.a.createElement("li",{className:"p-list__item is-ticked"},p("site-description2")),s.a.createElement("li",{className:"p-list__item is-ticked"},p("site-description3"))),s.a.createElement("p",null,p("build"))),s.a.createElement("div",null,s.a.createElement("a",{className:"p-button--brand",href:"/boards",alt:""},p("choose-board"))))))}}]),a}(n.Component),E=function(e){function a(){return Object(c.a)(this,a),Object(o.a)(this,Object(m.a)(a).apply(this,arguments))}return Object(u.a)(a,e),Object(i.a)(a,[{key:"renderBoards",value:function(){if(this.props.boards)return this.props.boards.map(function(e){return s.a.createElement("div",{className:"p-card board"},s.a.createElement("div",{className:"p-card__title"},s.a.createElement("img",{src:"/static/images/"+e.id+".png"})),s.a.createElement("div",{className:"p-card__content"},s.a.createElement("hr",null),s.a.createElement("h4",null,e.name),s.a.createElement("a",{href:"/boards/"+e.id,className:"p-button--neutral is-inline"},p("select-board"))))})}},{key:"render",value:function(){return s.a.createElement("div",{className:"first"},s.a.createElement("h2",null,p("choose-board")),s.a.createElement("p",null,p("choose-board-desc")),this.renderBoards())}}]),a}(n.Component),N=function(e){function a(){var e,t;Object(c.a)(this,a);for(var n=arguments.length,s=new Array(n),r=0;r<n;r++)s[r]=arguments[r];return(t=Object(o.a)(this,(e=Object(m.a)(a)).call.apply(e,[this].concat(s)))).handleClick=function(e){window.location.href="/boards/"+t.props.board.id+"/"+e.target.getAttribute("data-key")},t}return Object(u.a)(a,e),Object(i.a)(a,[{key:"renderOS",value:function(){var e=this;return this.props.board.os.map(function(a){var t=a.type+a.version;return s.a.createElement("div",{className:"os",key:t,"data-key":t,onClick:e.handleClick},s.a.createElement("img",{src:"/static/images/ubuntu.png",alt:"Ubuntu","data-key":t}),s.a.createElement("h4",{"data-key":t},p(a.type)),s.a.createElement("p",{"data-key":t},p(a.version)))})}},{key:"render",value:function(){var e="/static/images/"+this.props.board.id+".png";return s.a.createElement("div",{className:"first"},s.a.createElement("div",{className:"row"},s.a.createElement("div",{className:"col-1"},s.a.createElement("img",{src:e,alt:this.props.board.name,width:"100px"})),s.a.createElement("div",{className:"col-5 u-align--bottom"},s.a.createElement("h4",null,this.props.board.name))),s.a.createElement("h2",null,p("choose-os")),s.a.createElement("p",null,p("choose-os-desc")),this.renderOS())}}]),a}(n.Component),y=t(7),k=t.n(y),O="/v1/";var j={baseUrl:window.location.protocol+"//"+window.location.hostname+":"+window.location.port+O,LoadingImage:"/static/images/ajax-loader.gif"},C={boardsList:function(e,a){return k.a.get(j.baseUrl+"boards")},storeSearch:function(e,a){return k.a.get(j.baseUrl+"store/snaps/"+e)}},S=function(e){function a(e){var t;return Object(c.a)(this,a),(t=Object(o.a)(this,Object(m.a)(a).call(this,e))).pageUp=function(){var e=t.calculatePages(),a=t.state.page+1;a>e&&(a=e),t.setState({page:a}),t.signalPageChange(a)},t.pageDown=function(){var e=t.state.page-1;e<=0&&(e=1),t.setState({page:e}),t.signalPageChange(e)},t.state={page:1,query:null,maxRecords:5|e.pageSize},t}return Object(u.a)(a,e),Object(i.a)(a,[{key:"signalPageChange",value:function(e){var a=(e-1)*this.state.maxRecords;this.props.pageChange(a,a+this.state.maxRecords)}},{key:"calculatePages",value:function(){var e=this.props.displayRows.length,a=parseInt(e/this.state.maxRecords,10);return e%this.state.maxRecords>0&&(a+=1),a}},{key:"renderPaging",value:function(){var e=this.calculatePages();return e>1?s.a.createElement("div",{className:"u-float--right spacer"},s.a.createElement("button",{className:"p-button--neutral small",href:"",onClick:this.pageDown},"\xab"),s.a.createElement("span",null,"\xa0",this.state.page," of ",e,"\xa0"),s.a.createElement("button",{className:"p-button--neutral small",href:"",onClick:this.pageUp},"\xbb")):s.a.createElement("div",{className:"u-float--right"})}},{key:"render",value:function(){return s.a.createElement("div",{className:"col-12"},this.renderPaging())}}]),a}(n.Component),_=t(18),w=t.n(_),I=5,R=function(e){function a(e){var t;return Object(c.a)(this,a),(t=Object(o.a)(this,Object(m.a)(a).call(this,e))).handleClear=function(e){t.setState({snaps:[]})},t.handleSearchChange=function(e){e.preventDefault(),t.setState({snapName:e.target.value})},t.handleKeyPress=function(e){"Enter"===e.key&&t.handleSearchStore(e)},t.handleSearchStore=function(e){e.preventDefault(),t.searchStore()},t.handleInstall=function(e){e.preventDefault();var a=e.target.getAttribute("data-key");t.props.handleInstallClick(a)},t.handleRecordsForPage=function(e,a){t.setState({startRow:e,endRow:a})},t.state={snapName:"",snaps:[],loadingSearch:!1,page:1,startRow:0,endRow:I},t}return Object(u.a)(a,e),Object(i.a)(a,[{key:"searchStore",value:function(){var e=this;0!==this.state.snapName.length&&(this.setState({loadingSearch:!0}),C.storeSearch(this.state.snapName).then(function(a){a.data._embedded&&a.data._embedded["clickindex:package"]&&e.setState({snaps:a.data._embedded["clickindex:package"],loadingSearch:!1,message:null,messageType:null})}))}},{key:"renderSnaps",value:function(e){var a=this;return e.length>0?s.a.createElement("div",null,s.a.createElement("p",null,e.length," snaps found"),s.a.createElement("table",null,s.a.createElement("tbody",null,e.slice(this.state.startRow,this.state.endRow).map(function(e){return s.a.createElement("tr",{key:e.snap_id,title:e.description},s.a.createElement("td",{className:"small"},s.a.createElement("button",{"data-key":e.package_name,className:"p-button--neutral small",title:p("add-application"),onClick:a.handleInstall},p("add"))),s.a.createElement("td",{className:"overflow"},s.a.createElement("b",null,e.package_name)," ",e.version),s.a.createElement("td",{className:"overflow"},e.developer_name," (",w()(e.binary_filesize),")"))})))):s.a.createElement("div",null,"No snaps found.")}},{key:"render",value:function(){var e=this.state.snaps;return this.props.message?s.a.createElement("div",{className:"p-card col-6"},s.a.createElement("h3",{className:"p-card__title"},p("snap-store")),s.a.createElement("p",null,this.state.loadingSearch?s.a.createElement("img",{src:j.LoadingImage,alt:p("loading")}):"",this.props.message),s.a.createElement("p",null,s.a.createElement("form",{className:"p-search-box"},s.a.createElement("input",{className:"p-search-box__input",type:"search",name:"snapname",onKeyPress:this.handleKeyPress,onChange:this.handleSearchChange,placeholder:p("search-store")}),s.a.createElement("button",{type:"reset",className:"p-search-box__reset",alt:"reset",disabled:"",onClick:this.handleClear},s.a.createElement("i",{className:"p-icon--close"})),s.a.createElement("button",{type:"submit",onClick:this.handleSearchStore,className:"p-search-box__button",alt:"search"},s.a.createElement("i",{className:"p-icon--search"})))),s.a.createElement(S,{displayRows:e,pageSize:I,pageChange:this.handleRecordsForPage}),this.renderSnaps(e)):s.a.createElement("span",null)}}]),a}(n.Component),x=function(e){function a(e){var t;Object(c.a)(this,a),(t=Object(o.a)(this,Object(m.a)(a).call(this,e))).handleSnapInstall=function(e){if(!(t.state.snaps.indexOf(e)>=0)){var a=t.state.snaps;a.push(e),t.setState({snaps:a})}},t.handleSnapOnChange=function(e){e.preventDefault(),t.setState({snapName:e.target.value})},t.handleDialogCancel=function(e){e.preventDefault()},t.handleRemove=function(e){e.preventDefault();var a=e.target.getAttribute("data-key"),n=t.state.snaps.filter(function(e){return e!==a});t.setState({snaps:n})},t.handleConfirm=function(e){e.preventDefault(),function(e){var a=JSON.stringify(e);localStorage.setItem("board",a)}({board:t.props.board,os:t.props.os,snaps:t.state.snaps}),window.location.href="/confirm"},t.state={snaps:[]};var n=h();return n&&n.snaps&&n.snaps.length>0&&(t.state.snaps=n.snaps),t}return Object(u.a)(a,e),Object(i.a)(a,[{key:"renderSelected",value:function(){var e="/static/images/"+this.props.board.id+".png";return s.a.createElement("div",{className:"row"},s.a.createElement("div",{className:"col-1"},s.a.createElement("img",{src:e,alt:this.props.board.name,width:"100px"})),s.a.createElement("div",{className:"col-5 u-align--bottom"},s.a.createElement("h4",null,this.props.board.name),s.a.createElement("p",null,"Ubuntu ",p(this.props.os.type)," ",this.props.os.version)))}},{key:"render",value:function(){var e=this;return s.a.createElement("div",null,this.renderSelected(),s.a.createElement("h2",null,p("choose-applications")),s.a.createElement("p",null,p("choose-applications-desc")),s.a.createElement("div",{className:"row"},s.a.createElement("div",{className:"p-card col-4"},s.a.createElement("p",null,p("pre-installed-snaps")),s.a.createElement("ul",{className:"p-list"},this.state.snaps.map(function(a){return s.a.createElement("li",{className:"p-list__item"},s.a.createElement("button",{"data-key":a,className:"p-button--neutral small",title:p("remove-application"),onClick:e.handleRemove},p("remove")),"\xa0 ",a)}))),s.a.createElement(R,{message:p("find-snaps"),handleTextChange:this.handleSnapOnChange,handleInstallClick:this.handleSnapInstall,handleCancelClick:this.handleDialogCancel})),s.a.createElement("div",null,s.a.createElement("button",{className:"p-button--brand",onClick:this.handleConfirm},p("confirm"))))}}]),a}(n.Component),P=function(e){function a(e){var t;return Object(c.a)(this,a),(t=Object(o.a)(this,Object(m.a)(a).call(this,e))).handleBuild=function(e){setTimeout(function(){var e=["ogra","ondra","alfonso","konrad"],a=e[Math.floor(Math.random()*e.length)];t.setState({messages:["Waking trained chef... "+a]})},1e3),setTimeout(function(){var e=t.state.messages;e.push("Cranking the build machine..."),t.setState({messages:e})},2e3),setTimeout(function(){var e=t.state.messages;e.push("This is just a demo, but you get the idea"),t.setState({messages:e})},3e3)},t.state={messages:[]},t}return Object(u.a)(a,e),Object(i.a)(a,[{key:"renderSelected",value:function(){var e="/static/images/"+this.props.board.board.id+".png";return s.a.createElement("div",{className:"row"},s.a.createElement("div",{className:"col-1"},s.a.createElement("img",{src:e,alt:this.props.board.board.name,width:"100px"})),s.a.createElement("div",{className:"col-5 u-align--bottom"},s.a.createElement("h4",null,this.props.board.board.name),s.a.createElement("p",null,"Ubuntu ",p(this.props.board.os.type)," ",this.props.board.os.version),s.a.createElement("p",null,1!==this.props.board.snaps.length?this.props.board.snaps.length+" snaps":"1 snap")))}},{key:"renderConsole",value:function(){if(0!==this.state.messages.length)return s.a.createElement("pre",{className:"console"},this.state.messages.map(function(e){return e+"\n"}))}},{key:"render",value:function(){var e=this.props.board;return s.a.createElement("div",{className:"first"},this.renderSelected(),s.a.createElement("h2",null,p("confirm-configuration")),s.a.createElement("div",{className:"row"},s.a.createElement("pre",null,s.a.createElement("code",null,"board:","\n","  name: ",e.board.name,"\n","  os:","\n","    type: ",e.os.type,"\n","    version: ",e.os.version,"\n","  snaps:","\n",e.snaps.map(function(e){return"    "+e+"\n"})))),s.a.createElement("div",null,s.a.createElement("h3",null,p("ready-to-build")),s.a.createElement("button",{className:"p-button--brand",onClick:this.handleBuild},p("build"))),s.a.createElement("div",{className:"row"},this.renderConsole()))}}]),a}(n.Component),U=t(19),B=t.n(U),D=function(e){function a(){return Object(c.a)(this,a),Object(o.a)(this,Object(m.a)(a).apply(this,arguments))}return Object(u.a)(a,e),Object(i.a)(a,[{key:"render",value:function(){var e=this.props.route,a="boards"!==e.section||e.sectionId?"":"active",t="boards"===e.section&&e.sectionId&&!e.subsection?"active":"",n="boards"===e.section&&e.sectionId&&e.subsection?"active":"";if("confirm"===e.section){var r=h();console.log(r),e.section="boards",e.sectionId=r.board.id,e.subsection=r.os.type+r.os.version,e.subsubsection="confirm"}return s.a.createElement("div",{className:"breadcrumbs"},s.a.createElement("div",{className:"row"},s.a.createElement("ul",{className:"p-breadcrumbs"},s.a.createElement("li",{className:"p-breadcrumbs__item"},s.a.createElement("a",{href:"/"},p("home"))),s.a.createElement("li",{className:"p-breadcrumbs__item "+a},s.a.createElement("a",{href:"/boards"},p("boards"))),t||n||e.subsubsection?s.a.createElement("li",{className:"p-breadcrumbs__item "+t},s.a.createElement("a",{href:"/"+e.section+"/"+e.sectionId},p("os"))):"",n||e.subsubsection?s.a.createElement("li",{className:"p-breadcrumbs__item "+n},s.a.createElement("a",{href:"/"+e.section+"/"+e.sectionId+"/"+e.subsection},p("applications"))):"",e.subsubsection?s.a.createElement("li",{className:"p-breadcrumbs__item active"},s.a.createElement("a",{href:"/confirm"},p("confirm"))):"")))}}]),a}(n.Component),A=B()(),L=function(e){function a(e){var t;return Object(c.a)(this,a),(t=Object(o.a)(this,Object(m.a)(a).call(this,e))).state={location:A.location,token:e.token||{},boards:[]},t.getBoards(),t}return Object(u.a)(a,e),Object(i.a)(a,[{key:"getBoards",value:function(){var e=this;C.boardsList().then(function(a){e.setState({boards:a.data.boards})})}},{key:"renderBoards",value:function(e,a){if(!e)return s.a.createElement(E,{boards:this.state.boards});var t=this.state.boards.filter(function(a){return a.id===e});if(0!==t.length){if(!a)return s.a.createElement(N,{board:t[0]});var n=t[0].os.filter(function(e){return e.id===a});if(0!==n.length)return s.a.createElement(x,{board:t[0],os:n[0]})}}},{key:"renderConfirm",value:function(e,a){var t=h();return s.a.createElement(P,{board:t})}},{key:"render",value:function(){var e=function(){var e=window.location.pathname.split("/");switch(e.length){case 2:return{section:e[1]};case 3:return{section:e[1],sectionId:e[2]};case 4:return{section:e[1],sectionId:e[2],subsection:e[3]};default:return{}}}();return console.log(e),s.a.createElement("div",{className:"App"},""===e.section?s.a.createElement(g,null):"",""!==e.section?s.a.createElement(f,null):"",""!==e.section?s.a.createElement(D,{route:e}):"",s.a.createElement("div",{className:"content row"},""===e.section?s.a.createElement(v,null):"","boards"===e.section?this.renderBoards(e.sectionId,e.subsection):"","confirm"===e.section?this.renderConfirm(e.sectionId,e.subsection):""),s.a.createElement(b,null))}}]),a}(n.Component);l.a.render(s.a.createElement(L,null),document.getElementById("root"))}},[[20,1,2]]]);
//# sourceMappingURL=main.0ed2da23.chunk.js.map