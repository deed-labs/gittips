import{S as U,i as Y,s as z,k as m,U as C,a as M,q,l as p,m as d,V as N,h as c,c as O,r as P,n as t,b as H,L as s,O as L,D as E,M as S,o as A}from"../../../chunks/index-2fd92ac7.js";import{p as F}from"../../../chunks/stores-879647a7.js";import{s as G}from"../../../chunks/index-8d30d9a6.js";import{s as T}from"../../../chunks/storage-2e55a9a8.js";function J(v){let a,n,e,i,r,l,f,o,b,x,u,k,w,y;return{c(){a=m("div"),n=m("div"),e=m("div"),i=m("div"),r=C("svg"),l=C("path"),f=M(),o=m("h1"),b=q("You are all set!"),x=M(),u=m("button"),k=q("Close"),this.h()},l(_){a=p(_,"DIV",{class:!0});var g=d(a);n=p(g,"DIV",{class:!0});var I=d(n);e=p(I,"DIV",{class:!0});var h=d(e);i=p(h,"DIV",{class:!0});var D=d(i);r=N(D,"svg",{class:!0,fill:!0,viewBox:!0});var V=d(r);l=N(V,"path",{"stroke-linecap":!0,"stroke-linejoin":!0,"stroke-width":!0,d:!0}),d(l).forEach(c),V.forEach(c),D.forEach(c),f=O(h),o=p(h,"H1",{class:!0});var j=d(o);b=P(j,"You are all set!"),j.forEach(c),x=O(h),u=p(h,"BUTTON",{class:!0});var B=d(u);k=P(B,"Close"),B.forEach(c),h.forEach(c),I.forEach(c),g.forEach(c),this.h()},h(){t(l,"stroke-linecap","round"),t(l,"stroke-linejoin","round"),t(l,"stroke-width","2"),t(l,"d","M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"),t(r,"class","text-center stroke-success h-10 w-10"),t(r,"fill","none"),t(r,"viewBox","0 0 24 24"),t(i,"class","flex justify-center"),t(o,"class","text-4xl font-bold py-6"),t(u,"class","btn btn-primary"),t(e,"class","max-w-md"),t(n,"class","hero-content text-center"),t(a,"class","hero min-h-screen bg-base-200")},m(_,g){H(_,a,g),s(a,n),s(n,e),s(e,i),s(i,r),s(r,l),s(e,f),s(e,o),s(o,b),s(e,x),s(e,u),s(u,k),w||(y=L(u,"click",v[0]),w=!0)},p:E,i:E,o:E,d(_){_&&c(a),w=!1,y()}}}function K(v,a,n){let e,i;S(v,T,o=>n(1,e=o)),S(v,F,o=>n(2,i=o));let r=e.wallet_address,l=Number(i.url.searchParams.get("installation_id"))??0;return A(async()=>{await G(r,l),T.set({...e,bot_installation_done:!0})}),[()=>{window.close()}]}class Z extends U{constructor(a){super(),Y(this,a,K,J,z,{})}}export{Z as default};
