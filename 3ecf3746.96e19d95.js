(window.webpackJsonp=window.webpackJsonp||[]).push([[17],{141:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return i})),a.d(t,"metadata",(function(){return b})),a.d(t,"rightToc",(function(){return l})),a.d(t,"default",(function(){return d}));var r=a(1),n=a(9),c=(a(0),a(177)),i={sidebar_label:"Index",title:"Container Call [object]"},b={id:"opspec/reference/structure/op-directory/op/call/container/index",title:"Container Call [object]",description:"An object defining a container call.",source:"@site/docs/opspec/reference/structure/op-directory/op/call/container/index.md",permalink:"/docs/opspec/reference/structure/op-directory/op/call/container/index",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/call/container/index.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578713522,sidebar_label:"Index",sidebar:"docs",previous:{title:"Call [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/index"},next:{title:"Image [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/container/image"}},l=[{value:"Properties",id:"properties",children:[{value:"image",id:"image",children:[]},{value:"cmd",id:"cmd",children:[]},{value:"dirs",id:"dirs",children:[]},{value:"envVars",id:"envvars",children:[]},{value:"files",id:"files",children:[]},{value:"name",id:"name",children:[]},{value:"ports",id:"ports",children:[]},{value:"sockets",id:"sockets",children:[]},{value:"workDir",id:"workdir",children:[]}]}],o={rightToc:l},p="wrapper";function d(e){var t=e.components,a=Object(n.a)(e,["components"]);return Object(c.b)(p,Object(r.a)({},o,a,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An object defining a container call."),Object(c.b)("h2",{id:"properties"},"Properties"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"must have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#image"}),"image")))),Object(c.b)("li",{parentName:"ul"},"may have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#cmd"}),"cmd")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#dirs"}),"dirs")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#envvars"}),"envVars")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#files"}),"files")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#name"}),"name")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#ports"}),"ports")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#sockets"}),"sockets")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(r.a)({parentName:"li"},{href:"#workdir"}),"workDir"))))),Object(c.b)("h3",{id:"image"},"image"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/container/image"}),"image")," object defining the container image run by the call."),Object(c.b)("h3",{id:"cmd"},"cmd"),Object(c.b)("p",null,"An array of ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"../../../../../../types/string#initializer"}),"string initializers")," defining the path (from ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"#workdir"}),"workDir"),") of the binary to call and it's arguments."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"defining cmd overrides any entrypoint and/or cmd defined by the image")),Object(c.b)("h3",{id:"dirs"},"dirs"),Object(c.b)("p",null,"An object for which each key is an absolute path in the container and each value is one of:"),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(r.a)({parentName:"tr"},{align:null}),"value"),Object(c.b)("th",Object(r.a)({parentName:"tr"},{align:null}),"meaning"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),"Mount dir embedded in op w/ same path (equivalent to ",Object(c.b)("inlineCode",{parentName:"td"},"$(/absolute/path)"),")")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(r.a)({parentName:"td"},{href:"../../../../../../types/dir.md"}),"dir")," ",Object(c.b)("a",Object(r.a)({parentName:"td"},{href:"../../../reference.md"}),"reference")),Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),"Mount dir")))),Object(c.b)("h3",{id:"envvars"},"envVars"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"../../../../../../types/object#initializer"}),"object initializer")," or ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"../../../reference.md"}),"reference"),", whos properties represent the name and value of an environment variable to be set in the container."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"upon evaluation, the key and value of each property will be coerced to a string.")),Object(c.b)("h3",{id:"files"},"files"),Object(c.b)("p",null,"An object for which each key is an absolute path in the container and each value is one of:"),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(r.a)({parentName:"tr"},{align:null}),"value"),Object(c.b)("th",Object(r.a)({parentName:"tr"},{align:null}),"meaning"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),"Mount file embedded in op w/ same path (equivalent to ",Object(c.b)("inlineCode",{parentName:"td"},"$(/absolute/path)"),")")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(r.a)({parentName:"td"},{href:"../../../../../../types/file.md"}),"file")," ",Object(c.b)("a",Object(r.a)({parentName:"td"},{href:"../../../reference.md"}),"reference")),Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),"Mount file")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(r.a)({parentName:"td"},{href:"../../../../../../types/file#initializer"}),"file initializer")),Object(c.b)("td",Object(r.a)({parentName:"tr"},{align:null}),"Evaluate and mount")))),Object(c.b)("h3",{id:"name"},"name"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"../../../../../../types/string#initializer"}),"string initializer")," defining a name by which the container can be resolved on the opctl network."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"if multiple containers are given the same name, network requests will be distributed (load balanced) across them. ")),Object(c.b)("h3",{id:"ports"},"ports"),Object(c.b)("p",null,"An object defining container ports exposed on the opctl host where:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"each key is a container port or range of ports (optionally including protocol) matching ",Object(c.b)("inlineCode",{parentName:"li"},"[0-9]+(-[0-9]+)?(tcp|udp)")),Object(c.b)("li",{parentName:"ul"},"each value is a corresponding opctl host port or range of ports matching ",Object(c.b)("inlineCode",{parentName:"li"},"[0-9]+(-[0-9]+)?"))),Object(c.b)("h3",{id:"sockets"},"sockets"),Object(c.b)("p",null,"An object for which each key is an absolute path in the container and and each value is a ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"../../../../../../types/socket.md"}),"socket")," ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"../../../reference.md"}),"reference")," to mount. "),Object(c.b)("h3",{id:"workdir"},"workDir"),Object(c.b)("p",null,"An absolute path which defines where ",Object(c.b)("a",Object(r.a)({parentName:"p"},{href:"#cmd"}),"cmd")," is executed."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"defining workDir overrides any defined by the image")))}d.isMDXComponent=!0},177:function(e,t,a){"use strict";a.d(t,"a",(function(){return d})),a.d(t,"b",(function(){return s}));var r=a(0),n=a.n(r);function c(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function i(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,r)}return a}function b(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?i(Object(a),!0).forEach((function(t){c(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):i(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function l(e,t){if(null==e)return{};var a,r,n=function(e,t){if(null==e)return{};var a,r,n={},c=Object.keys(e);for(r=0;r<c.length;r++)a=c[r],t.indexOf(a)>=0||(n[a]=e[a]);return n}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(r=0;r<c.length;r++)a=c[r],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(n[a]=e[a])}return n}var o=n.a.createContext({}),p=function(e){var t=n.a.useContext(o),a=t;return e&&(a="function"==typeof e?e(t):b({},t,{},e)),a},d=function(e){var t=p(e.components);return n.a.createElement(o.Provider,{value:t},e.children)},u="mdxType",j={inlineCode:"code",wrapper:function(e){var t=e.children;return n.a.createElement(n.a.Fragment,{},t)}},m=Object(r.forwardRef)((function(e,t){var a=e.components,r=e.mdxType,c=e.originalType,i=e.parentName,o=l(e,["components","mdxType","originalType","parentName"]),d=p(a),u=r,m=d["".concat(i,".").concat(u)]||d[u]||j[u]||c;return a?n.a.createElement(m,b({ref:t},o,{components:a})):n.a.createElement(m,b({ref:t},o))}));function s(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var c=a.length,i=new Array(c);i[0]=m;var b={};for(var l in t)hasOwnProperty.call(t,l)&&(b[l]=t[l]);b.originalType=e,b[u]="string"==typeof e?e:r,i[1]=b;for(var o=2;o<c;o++)i[o]=a[o];return n.a.createElement.apply(null,i)}return n.a.createElement.apply(null,a)}m.displayName="MDXCreateElement"}}]);