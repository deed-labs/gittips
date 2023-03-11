import{S as Se,i as je,s as ze,x as Ae,a as H,k as d,y as Be,c as O,l as h,m,h as l,n as s,z as Te,b as L,L as t,N as se,O as re,f as C,t as Fe,A as Ne,F as Ue,o as Qe,q as R,r as G,P as J,Q as K,D as X,R as ve}from"../../../chunks/index-2fd92ac7.js";import{H as We,f as ee,a as me}from"../../../chunks/Header-42c274a2.js";const qe=""+new URL("../../../assets/issue_screenshot-5f3d4fab.png",import.meta.url).href,Je=""+new URL("../../../assets/comment_screenshot-f32c8ad2.png",import.meta.url).href,Ke=""+new URL("../../../assets/payment_screenshot-692a51e7.png",import.meta.url).href;function Xe(n){return n()}function Ye(n){n.forEach(Xe)}const oe=[],Ve=[],fe=[],De=[],Me=Promise.resolve();let pe=!1;function Ze(){pe||(pe=!0,Me.then(tt))}function $e(){return Ze(),Me}function et(n){fe.push(n)}const he=new Set;let ce=0;function tt(){do{for(;ce<oe.length;){const n=oe[ce];ce++,lt(n.$$)}for(oe.length=0,ce=0;Ve.length;)Ve.pop()();for(let n=0;n<fe.length;n+=1){const e=fe[n];he.has(e)||(he.add(e),e())}fe.length=0}while(oe.length);for(;De.length;)De.pop()();pe=!1,he.clear()}function lt(n){if(n.fragment!==null){n.update(),Ye(n.before_update);const e=n.dirty;n.dirty=[-1],n.fragment&&n.fragment.p(n.ctx,e),n.after_update.forEach(et)}}const st={root:null,rootMargin:"0px",threshold:0,unobserveOnEnter:!1},ue=(n,e)=>new CustomEvent(n,{detail:e});function ae(n,e={}){const{root:r,rootMargin:x,threshold:o,unobserveOnEnter:a}=Object.assign(Object.assign({},st),e);let c={x:void 0,y:void 0},b={vertical:void 0,horizontal:void 0};if(typeof IntersectionObserver<"u"&&n){const p=new IntersectionObserver((v,u)=>{v.forEach(i=>{c.y>i.boundingClientRect.y?b.vertical="up":b.vertical="down",c.x>i.boundingClientRect.x?b.horizontal="left":b.horizontal="right",c={y:i.boundingClientRect.y,x:i.boundingClientRect.x};const _={inView:i.isIntersecting,entry:i,scrollDirection:b,node:n,observer:u};n.dispatchEvent(ue("change",_)),i.isIntersecting?(n.dispatchEvent(ue("enter",_)),a&&u.unobserve(n)):n.dispatchEvent(ue("leave",_))})},{root:r,rootMargin:x,threshold:o});return $e().then(()=>{n.dispatchEvent(ue("init",{observer:p,node:n}))}),p.observe(n),{destroy(){p.unobserve(n)}}}}function He(n){let e,r,x,o,a,c,b,p,v,u,i,_,y,E;return{c(){e=d("div"),r=d("h1"),x=R("Reward the community, fuel innovation"),o=H(),a=d("p"),c=R("Gittips is a bot for GitHub that enables owners of open-source projects to reward contributors with cryptocurrency"),b=H(),p=d("div"),v=d("a"),u=R("Docs"),i=H(),_=d("a"),y=R("Explore"),this.h()},l(w){e=h(w,"DIV",{class:!0});var g=m(e);r=h(g,"H1",{class:!0});var k=m(r);x=G(k,"Reward the community, fuel innovation"),k.forEach(l),o=O(g),a=h(g,"P",{class:!0});var D=m(a);c=G(D,"Gittips is a bot for GitHub that enables owners of open-source projects to reward contributors with cryptocurrency"),D.forEach(l),b=O(g),p=h(g,"DIV",{class:!0});var F=m(p);v=h(F,"A",{class:!0,target:!0,rel:!0,href:!0});var P=m(v);u=G(P,"Docs"),P.forEach(l),i=O(F),_=h(F,"A",{class:!0,href:!0});var I=m(_);y=G(I,"Explore"),I.forEach(l),F.forEach(l),g.forEach(l),this.h()},h(){s(r,"class","text-7xl font-bold text-white"),s(a,"class","text-xl py-10 text-gray-400"),s(v,"class","btn btn-wide btn-primary btn-outline text-white font-bold rounded-full capitalize"),s(v,"target","_blank"),s(v,"rel","noreferrer"),s(v,"href","https://deed-labs.gitbook.io/gittips/"),s(_,"class","btn btn-wide btn-primary text-white font-bold rounded-full capitalize"),s(_,"href","/explore/bounties"),s(p,"class","flex flex-col md:flex-row items-center justify-center gap-4"),s(e,"class","max-w-4xl")},m(w,g){L(w,e,g),t(e,r),t(r,x),t(e,o),t(e,a),t(a,c),t(e,b),t(e,p),t(p,v),t(v,u),t(p,i),t(p,_),t(_,y)},i(w){E||J(()=>{E=K(e,ee,{y:200,duration:1500}),E.start()})},o:X,d(w){w&&l(e)}}}function Oe(n){let e,r,x,o,a,c,b,p,v;return{c(){e=d("div"),r=d("h2"),x=R("Our solution does not interfere with your "),o=d("span"),a=R("usual workflow"),c=H(),b=d("p"),p=R("Create tasks and send payments without leaving GitHub"),this.h()},l(u){e=h(u,"DIV",{});var i=m(e);r=h(i,"H2",{class:!0});var _=m(r);x=G(_,"Our solution does not interfere with your "),o=h(_,"SPAN",{class:!0});var y=m(o);a=G(y,"usual workflow"),y.forEach(l),_.forEach(l),c=O(i),b=h(i,"P",{class:!0});var E=m(b);p=G(E,"Create tasks and send payments without leaving GitHub"),E.forEach(l),i.forEach(l),this.h()},h(){s(o,"class","text-primary"),s(r,"class","text-4xl text-white font-bold"),s(b,"class","text-2xl text-gray-400 mt-5")},m(u,i){L(u,e,i),t(e,r),t(r,x),t(r,o),t(o,a),t(e,c),t(e,b),t(b,p)},i(u){v||J(()=>{v=K(e,me,{duration:1500}),v.start()})},o:X,d(u){u&&l(e)}}}function Re(n){let e,r,x,o,a,c,b,p,v,u,i,_,y,E,w,g,k,D,F,P;return{c(){e=d("div"),r=d("h1"),x=R("Automated Bounty Creation"),o=H(),a=d("p"),c=R("We automate the process of creating bounties from issues on GitHub. "),b=d("br"),p=R(" Find developers for tasks without breaking your usual flow."),v=H(),u=d("p"),i=R("Just add "),_=d("span"),y=R("bounty"),E=R(" label to an issue."),g=H(),k=d("div"),D=d("img"),this.h()},l(I){e=h(I,"DIV",{class:!0});var N=m(e);r=h(N,"H1",{class:!0});var W=m(r);x=G(W,"Automated Bounty Creation"),W.forEach(l),o=O(N),a=h(N,"P",{class:!0});var Y=m(a);c=G(Y,"We automate the process of creating bounties from issues on GitHub. "),b=h(Y,"BR",{}),p=G(Y," Find developers for tasks without breaking your usual flow."),Y.forEach(l),v=O(N),u=h(N,"P",{class:!0});var U=m(u);i=G(U,"Just add "),_=h(U,"SPAN",{class:!0});var Z=m(_);y=G(Z,"bounty"),Z.forEach(l),E=G(U," label to an issue."),U.forEach(l),N.forEach(l),g=O(I),k=h(I,"DIV",{class:!0});var Q=m(k);D=h(Q,"IMG",{class:!0,src:!0,alt:!0}),Q.forEach(l),this.h()},h(){s(r,"class","text-4xl font-bold"),s(a,"class","text-xl text-gray-400 mt-5"),s(_,"class","badge badge-primary badge-outline"),s(u,"class","text-xl text-gray-400"),s(e,"class","w-full md:w-1/2"),s(D,"class","card border border-primary image-full w-full md:w-5/6"),ve(D.src,F=qe)||s(D,"src",F),s(D,"alt","issue example"),s(k,"class","flex justify-center w-full md:w-1/2")},m(I,N){L(I,e,N),t(e,r),t(r,x),t(e,o),t(e,a),t(a,c),t(a,b),t(a,p),t(e,v),t(e,u),t(u,i),t(u,_),t(_,y),t(u,E),L(I,g,N),L(I,k,N),t(k,D)},p:X,i(I){w||J(()=>{w=K(e,ee,{x:-200,duration:1e3}),w.start()}),P||J(()=>{P=K(k,ee,{x:200,duration:1e3}),P.start()})},o:X,d(I){I&&l(e),I&&l(g),I&&l(k)}}}function Ge(n){let e,r,x,o,a,c,b,p,v,u,i,_;return{c(){e=d("div"),r=d("img"),a=H(),c=d("div"),b=d("h1"),p=R("Budget Control"),v=H(),u=d("p"),i=R("Each organization and user has its own balance reserved on the smart contract. So they can't accidentally spend more than planned."),this.h()},l(y){e=h(y,"DIV",{class:!0});var E=m(e);r=h(E,"IMG",{class:!0,src:!0,alt:!0}),E.forEach(l),a=O(y),c=h(y,"DIV",{class:!0});var w=m(c);b=h(w,"H1",{class:!0});var g=m(b);p=G(g,"Budget Control"),g.forEach(l),v=O(w),u=h(w,"P",{class:!0});var k=m(u);i=G(k,"Each organization and user has its own balance reserved on the smart contract. So they can't accidentally spend more than planned."),k.forEach(l),w.forEach(l),this.h()},h(){s(r,"class","card border border-primary w-full md:w-5/6"),ve(r.src,x=Je)||s(r,"src",x),s(r,"alt","issue example"),s(e,"class","w-full md:w-1/2"),s(b,"class","text-4xl font-bold"),s(u,"class","text-xl text-gray-400 mt-5"),s(c,"class","w-full md:w-1/2")},m(y,E){L(y,e,E),t(e,r),L(y,a,E),L(y,c,E),t(c,b),t(b,p),t(c,v),t(c,u),t(u,i)},p:X,i(y){o||J(()=>{o=K(e,ee,{x:-200,duration:1e3}),o.start()}),_||J(()=>{_=K(c,ee,{x:200,duration:1e3}),_.start()})},o:X,d(y){y&&l(e),y&&l(a),y&&l(c)}}}function Le(n){let e,r,x,o,a,c,b,p,v,u,i,_,y,E;return{c(){e=d("div"),r=d("h1"),x=R("Wallet Integration"),o=H(),a=d("p"),c=R("Gittips integrates with blockchain, allowing contributors to receive their rewards directly in their wallets. "),b=d("br"),p=R(" This makes it easy for contributors to use rewards as they see fit."),u=H(),i=d("div"),_=d("img"),this.h()},l(w){e=h(w,"DIV",{class:!0});var g=m(e);r=h(g,"H1",{class:!0});var k=m(r);x=G(k,"Wallet Integration"),k.forEach(l),o=O(g),a=h(g,"P",{class:!0});var D=m(a);c=G(D,"Gittips integrates with blockchain, allowing contributors to receive their rewards directly in their wallets. "),b=h(D,"BR",{}),p=G(D," This makes it easy for contributors to use rewards as they see fit."),D.forEach(l),g.forEach(l),u=O(w),i=h(w,"DIV",{class:!0});var F=m(i);_=h(F,"IMG",{class:!0,src:!0,alt:!0}),F.forEach(l),this.h()},h(){s(r,"class","text-4xl font-bold"),s(a,"class","text-xl text-gray-400 mt-5"),s(e,"class","w-full md:w-1/2"),s(_,"class","card border border-primary image-full w-full md:w-5/6"),ve(_.src,y=Ke)||s(_,"src",y),s(_,"alt","issue example"),s(i,"class","flex justify-center w-full md:w-1/2")},m(w,g){L(w,e,g),t(e,r),t(r,x),t(e,o),t(e,a),t(a,c),t(a,b),t(a,p),L(w,u,g),L(w,i,g),t(i,_)},p:X,i(w){v||J(()=>{v=K(e,ee,{x:-200,duration:1e3}),v.start()}),E||J(()=>{E=K(i,ee,{x:200,duration:1e3}),E.start()})},o:X,d(w){w&&l(e),w&&l(u),w&&l(i)}}}function Ce(n){let e,r,x,o,a,c,b;return{c(){e=d("div"),r=d("h1"),x=R("Empower your developer community"),o=H(),a=d("a"),c=R("Try it out"),this.h()},l(p){e=h(p,"DIV",{});var v=m(e);r=h(v,"H1",{class:!0});var u=m(r);x=G(u,"Empower your developer community"),u.forEach(l),o=O(v),a=h(v,"A",{class:!0,target:!0,rel:!0,href:!0});var i=m(a);c=G(i,"Try it out"),i.forEach(l),v.forEach(l),this.h()},h(){s(r,"class","text-white text-4xl"),s(a,"class","btn btn-wide btn-primary text-white font-bold rounded-full mt-12"),s(a,"target","_blank"),s(a,"rel","noreferrer"),s(a,"href","https://deed-labs.gitbook.io/gittips/")},m(p,v){L(p,e,v),t(e,r),t(r,x),t(e,o),t(e,a),t(a,c)},i(p){b||J(()=>{b=K(e,me,{duration:1500}),b.start()})},o:X,d(p){p&&l(e)}}}function Pe(n){let e,r,x,o,a,c,b,p,v,u,i,_,y,E;return{c(){e=d("div"),r=d("h1"),x=R("Roadmap 2023"),o=H(),a=d("ul"),c=d("li"),b=R("Launch"),p=H(),v=d("li"),u=R("NFT Rewards Support"),i=H(),_=d("li"),y=R("Deeper GitHub integration"),this.h()},l(w){e=h(w,"DIV",{class:!0});var g=m(e);r=h(g,"H1",{class:!0});var k=m(r);x=G(k,"Roadmap 2023"),k.forEach(l),o=O(g),a=h(g,"UL",{class:!0});var D=m(a);c=h(D,"LI",{"data-content":!0,class:!0});var F=m(c);b=G(F,"Launch"),F.forEach(l),p=O(D),v=h(D,"LI",{"data-content":!0,class:!0});var P=m(v);u=G(P,"NFT Rewards Support"),P.forEach(l),i=O(D),_=h(D,"LI",{"data-content":!0,class:!0});var I=m(_);y=G(I,"Deeper GitHub integration"),I.forEach(l),D.forEach(l),g.forEach(l),this.h()},h(){s(r,"class","text-white text-4xl"),s(c,"data-content","Q1"),s(c,"class","step step-primary"),s(v,"data-content","Q1"),s(v,"class","step"),s(_,"data-content","Q2"),s(_,"class","step"),s(a,"class","steps steps-vertical md:steps-horizontal mt-12 text-white"),s(e,"class","card bg-gray-700 w-full md:w-2/3 text-center p-12")},m(w,g){L(w,e,g),t(e,r),t(r,x),t(e,o),t(e,a),t(a,c),t(c,b),t(a,p),t(a,v),t(v,u),t(a,i),t(a,_),t(_,y)},i(w){E||J(()=>{E=K(e,me,{duration:1500}),E.start()})},o:X,d(w){w&&l(e)}}}function rt(n){let e,r,x,o,a,c,b,p,v,u,i,_,y,E,w,g,k,D,F,P,I,N,W,Y,U,Z,Q,ie,q,$,de,_e;e=new We({props:{class:"bg-base-100"}});let z=n[0]&&He(),A=n[1][0]&&Oe(),M=n[1][1]&&Re(),S=n[1][2]&&Ge(),j=n[1][3]&&Le(),B=n[1][4]&&Ce(),T=n[1][5]&&Pe();return{c(){Ae(e.$$.fragment),r=H(),x=d("div"),o=d("div"),z&&z.c(),a=H(),c=d("div"),A&&A.c(),b=H(),p=d("div"),v=d("div"),u=d("ul"),i=d("li"),y=H(),E=d("li"),g=H(),k=d("li"),F=H(),P=d("div"),I=d("div"),M&&M.c(),N=H(),W=d("div"),S&&S.c(),Y=H(),U=d("div"),j&&j.c(),Z=H(),Q=d("div"),B&&B.c(),ie=H(),q=d("div"),T&&T.c(),this.h()},l(f){Be(e.$$.fragment,f),r=O(f),x=h(f,"DIV",{class:!0});var V=m(x);o=h(V,"DIV",{class:!0});var be=m(o);z&&z.l(be),be.forEach(l),V.forEach(l),a=O(f),c=h(f,"DIV",{class:!0});var we=m(c);A&&A.l(we),we.forEach(l),b=O(f),p=h(f,"DIV",{class:!0});var ne=m(p);v=h(ne,"DIV",{class:!0});var xe=m(v);u=h(xe,"UL",{class:!0});var te=m(u);i=h(te,"LI",{"data-content":!0,class:!0}),m(i).forEach(l),y=O(te),E=h(te,"LI",{"data-content":!0,class:!0}),m(E).forEach(l),g=O(te),k=h(te,"LI",{"data-content":!0,class:!0}),m(k).forEach(l),te.forEach(l),xe.forEach(l),F=O(ne),P=h(ne,"DIV",{class:!0});var le=m(P);I=h(le,"DIV",{class:!0});var ge=m(I);M&&M.l(ge),ge.forEach(l),N=O(le),W=h(le,"DIV",{class:!0});var ye=m(W);S&&S.l(ye),ye.forEach(l),Y=O(le),U=h(le,"DIV",{class:!0});var Ee=m(U);j&&j.l(Ee),Ee.forEach(l),le.forEach(l),ne.forEach(l),Z=O(f),Q=h(f,"DIV",{class:!0});var ke=m(Q);B&&B.l(ke),ke.forEach(l),ie=O(f),q=h(f,"DIV",{class:!0});var Ie=m(q);T&&T.l(Ie),Ie.forEach(l),this.h()},h(){s(o,"class","hero-content text-center"),s(x,"class","hero min-h-screen"),s(c,"class","flex flex-col items-center text-center justify-center p-28"),s(i,"data-content","✨"),s(i,"class",_="step "+(n[1][1]?"step-primary":"")),s(E,"data-content","💰"),s(E,"class",w="step "+(n[1][2]?"step-primary":"")),s(k,"data-content","💎"),s(k,"class",D="step "+(n[1][3]?"step-primary":"")),s(u,"class","steps steps-vertical h-full"),s(v,"class","flex flex-col justify-center items-center w-2/12 md:w-1/12"),s(I,"class","flex flex-col md:flex-row items-center gap-5 p-12 md:h-96"),s(W,"class","flex flex-col-reverse md:flex-row items-center gap-5 p-12 md:h-96"),s(U,"class","flex flex-col md:flex-row items-center gap-5 p-12 h-full md:h-96"),s(P,"class","w-10/12 md:w-11/12 h-full"),s(p,"class","flex flex-row items-stretch py-12 max-w-screen-2xl m-auto text-white"),s(Q,"class","py-24 text-center"),s(q,"class","flex items-center justify-center max-w-screen-2xl w-full h-96 m-auto py-12")},m(f,V){Te(e,f,V),L(f,r,V),L(f,x,V),t(x,o),z&&z.m(o,null),L(f,a,V),L(f,c,V),A&&A.m(c,null),L(f,b,V),L(f,p,V),t(p,v),t(v,u),t(u,i),t(u,y),t(u,E),t(u,g),t(u,k),t(p,F),t(p,P),t(P,I),M&&M.m(I,null),t(P,N),t(P,W),S&&S.m(W,null),t(P,Y),t(P,U),j&&j.m(U,null),L(f,Z,V),L(f,Q,V),B&&B.m(Q,null),L(f,ie,V),L(f,q,V),T&&T.m(q,null),$=!0,de||(_e=[se(ae.call(null,c,{unobserveOnEnter:!0,rootMargin:"-20%"})),re(c,"change",n[2]),se(ae.call(null,I,{unobserveOnEnter:!0,rootMargin:"-20%"})),re(I,"change",n[3]),se(ae.call(null,W,{unobserveOnEnter:!0,rootMargin:"-20%"})),re(W,"change",n[4]),se(ae.call(null,U,{unobserveOnEnter:!0,rootMargin:"-20%"})),re(U,"change",n[5]),se(ae.call(null,Q,{unobserveOnEnter:!0,rootMargin:"-20%"})),re(Q,"change",n[6]),se(ae.call(null,q,{unobserveOnEnter:!0,rootMargin:"-20%"})),re(q,"change",n[7])],de=!0)},p(f,[V]){f[0]?z?V&1&&C(z,1):(z=He(),z.c(),C(z,1),z.m(o,null)):z&&(z.d(1),z=null),f[1][0]?A?V&2&&C(A,1):(A=Oe(),A.c(),C(A,1),A.m(c,null)):A&&(A.d(1),A=null),(!$||V&2&&_!==(_="step "+(f[1][1]?"step-primary":"")))&&s(i,"class",_),(!$||V&2&&w!==(w="step "+(f[1][2]?"step-primary":"")))&&s(E,"class",w),(!$||V&2&&D!==(D="step "+(f[1][3]?"step-primary":"")))&&s(k,"class",D),f[1][1]?M?(M.p(f,V),V&2&&C(M,1)):(M=Re(),M.c(),C(M,1),M.m(I,null)):M&&(M.d(1),M=null),f[1][2]?S?(S.p(f,V),V&2&&C(S,1)):(S=Ge(),S.c(),C(S,1),S.m(W,null)):S&&(S.d(1),S=null),f[1][3]?j?(j.p(f,V),V&2&&C(j,1)):(j=Le(),j.c(),C(j,1),j.m(U,null)):j&&(j.d(1),j=null),f[1][4]?B?V&2&&C(B,1):(B=Ce(),B.c(),C(B,1),B.m(Q,null)):B&&(B.d(1),B=null),f[1][5]?T?V&2&&C(T,1):(T=Pe(),T.c(),C(T,1),T.m(q,null)):T&&(T.d(1),T=null)},i(f){$||(C(e.$$.fragment,f),C(z),C(A),C(M),C(S),C(j),C(B),C(T),$=!0)},o(f){Fe(e.$$.fragment,f),$=!1},d(f){Ne(e,f),f&&l(r),f&&l(x),z&&z.d(),f&&l(a),f&&l(c),A&&A.d(),f&&l(b),f&&l(p),M&&M.d(),S&&S.d(),j&&j.d(),f&&l(Z),f&&l(Q),B&&B.d(),f&&l(ie),f&&l(q),T&&T.d(),de=!1,Ue(_e)}}}function at(n,e,r){let x=!1;Qe(()=>r(0,x=!0));let o=new Array(6);return[x,o,({detail:i})=>{r(1,o[0]=i.inView,o)},({detail:i})=>{r(1,o[1]=i.inView,o)},({detail:i})=>{r(1,o[2]=i.inView,o)},({detail:i})=>{r(1,o[3]=i.inView,o)},({detail:i})=>{r(1,o[4]=i.inView,o)},({detail:i})=>{r(1,o[5]=i.inView,o)}]}class ot extends Se{constructor(e){super(),je(this,e,at,rt,ze,{})}}export{ot as default};
