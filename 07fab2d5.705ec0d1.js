(window.webpackJsonp=window.webpackJsonp||[]).push([[4],{123:function(e,r,t){"use strict";t.r(r),t.d(r,"frontMatter",(function(){return c})),t.d(r,"metadata",(function(){return o})),t.d(r,"rightToc",(function(){return i})),t.d(r,"default",(function(){return u}));var a=t(1),n=t(9),l=(t(0),t(177)),c={sidebar_label:"Index",title:"Call [object]"},o={id:"opspec/reference/structure/op-directory/op/call/index",title:"Call [object]",description:"An object defining a call in an operations call graph.",source:"@site/docs/opspec/reference/structure/op-directory/op/call/index.md",permalink:"/docs/opspec/reference/structure/op-directory/op/call/index",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/call/index.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578713522,sidebar_label:"Index",sidebar:"docs",previous:{title:"Op [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/index"},next:{title:"Container Call [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/container/index"}},i=[{value:"Properties",id:"properties",children:[{value:"container",id:"container",children:[]},{value:"op",id:"op",children:[]},{value:"parallel",id:"parallel",children:[]},{value:"parallelLoop",id:"parallelloop",children:[]},{value:"serial",id:"serial",children:[]},{value:"serialLoop",id:"serialloop",children:[]},{value:"if",id:"if",children:[]}]}],p={rightToc:i},b="wrapper";function u(e){var r=e.components,t=Object(n.a)(e,["components"]);return Object(l.b)(b,Object(a.a)({},p,t,{components:r,mdxType:"MDXLayout"}),Object(l.b)("p",null,"An object defining a call in an operations call graph."),Object(l.b)("h2",{id:"properties"},"Properties"),Object(l.b)("ul",null,Object(l.b)("li",{parentName:"ul"},"must have exactly one of",Object(l.b)("ul",{parentName:"li"},Object(l.b)("li",{parentName:"ul"},Object(l.b)("a",Object(a.a)({parentName:"li"},{href:"#container"}),"container")),Object(l.b)("li",{parentName:"ul"},Object(l.b)("a",Object(a.a)({parentName:"li"},{href:"#op"}),"op")),Object(l.b)("li",{parentName:"ul"},Object(l.b)("a",Object(a.a)({parentName:"li"},{href:"#parallel"}),"parallel")),Object(l.b)("li",{parentName:"ul"},Object(l.b)("a",Object(a.a)({parentName:"li"},{href:"#parallelloop"}),"parallelLoop")),Object(l.b)("li",{parentName:"ul"},Object(l.b)("a",Object(a.a)({parentName:"li"},{href:"#serial"}),"serial")),Object(l.b)("li",{parentName:"ul"},Object(l.b)("a",Object(a.a)({parentName:"li"},{href:"#serialloop"}),"serialLoop")))),Object(l.b)("li",{parentName:"ul"},"may have",Object(l.b)("ul",{parentName:"li"},Object(l.b)("li",{parentName:"ul"},Object(l.b)("a",Object(a.a)({parentName:"li"},{href:"#if"}),"if"))))),Object(l.b)("h3",{id:"container"},"container"),Object(l.b)("p",null,"A ",Object(l.b)("a",Object(a.a)({parentName:"p"},{href:"container/index"}),"container-call")," object defining a container to run."),Object(l.b)("h3",{id:"op"},"op"),Object(l.b)("p",null,"An ",Object(l.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/op"}),"op-call")," object defining an op to run."),Object(l.b)("h3",{id:"parallel"},"parallel"),Object(l.b)("p",null,"An array of ",Object(l.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/index"}),"call")," objects defining calls run in parallel (all at once without order)."),Object(l.b)("h3",{id:"parallelloop"},"parallelLoop"),Object(l.b)("p",null,"A ",Object(l.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/parallel-loop"}),"parallel-loop-call")," object defining a call loop in which all iterations happen in parallel (all at once without order)."),Object(l.b)("h3",{id:"serial"},"serial"),Object(l.b)("p",null,"An array of ",Object(l.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/index"}),"call")," objects defining calls run in serial (one after another in order)."),Object(l.b)("h3",{id:"serialloop"},"serialLoop"),Object(l.b)("p",null,"A ",Object(l.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/serial-loop"}),"serial-loop-call")," object defining a call loop in which each iteration happens in serial (one after another in order)"),Object(l.b)("h3",{id:"if"},"if"),Object(l.b)("p",null,"An array of ",Object(l.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/predicate"}),"predicates")," which must all be true for the call to take place."))}u.isMDXComponent=!0},177:function(e,r,t){"use strict";t.d(r,"a",(function(){return u})),t.d(r,"b",(function(){return j}));var a=t(0),n=t.n(a);function l(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function c(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);r&&(a=a.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,a)}return t}function o(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?c(Object(t),!0).forEach((function(r){l(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):c(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function i(e,r){if(null==e)return{};var t,a,n=function(e,r){if(null==e)return{};var t,a,n={},l=Object.keys(e);for(a=0;a<l.length;a++)t=l[a],r.indexOf(t)>=0||(n[t]=e[t]);return n}(e,r);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(a=0;a<l.length;a++)t=l[a],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(n[t]=e[t])}return n}var p=n.a.createContext({}),b=function(e){var r=n.a.useContext(p),t=r;return e&&(t="function"==typeof e?e(r):o({},r,{},e)),t},u=function(e){var r=b(e.components);return n.a.createElement(p.Provider,{value:r},e.children)},s="mdxType",d={inlineCode:"code",wrapper:function(e){var r=e.children;return n.a.createElement(n.a.Fragment,{},r)}},f=Object(a.forwardRef)((function(e,r){var t=e.components,a=e.mdxType,l=e.originalType,c=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),u=b(t),s=a,f=u["".concat(c,".").concat(s)]||u[s]||d[s]||l;return t?n.a.createElement(f,o({ref:r},p,{components:t})):n.a.createElement(f,o({ref:r},p))}));function j(e,r){var t=arguments,a=r&&r.mdxType;if("string"==typeof e||a){var l=t.length,c=new Array(l);c[0]=f;var o={};for(var i in r)hasOwnProperty.call(r,i)&&(o[i]=r[i]);o.originalType=e,o[s]="string"==typeof e?e:a,c[1]=o;for(var p=2;p<l;p++)c[p]=t[p];return n.a.createElement.apply(null,c)}return n.a.createElement.apply(null,t)}f.displayName="MDXCreateElement"}}]);