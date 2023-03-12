import{S as Se,i as je,s as ze,x as Ae,a as O,k as u,y as We,c as P,l as f,m as _,h as l,n as s,z as Be,b as L,L as t,N as le,O as se,f as C,t as Te,A as Fe,F as Ne,o as Ue,q as R,r as G,P as q,Q as J,D as K,R as pe}from"../../../chunks/index-2fd92ac7.js";import{H as Qe,f as Me,a as me}from"../../../chunks/Header-81c39ffe.js";const qe=""+new URL("../../../assets/issue_screenshot-5f3d4fab.png",import.meta.url).href,Je=""+new URL("../../../assets/comment_screenshot-f32c8ad2.png",import.meta.url).href,Ke=""+new URL("../../../assets/payment_screenshot-692a51e7.png",import.meta.url).href;function Xe(r){return r()}function Ye(r){r.forEach(Xe)}const ne=[],ke=[],ue=[],Ve=[],Ce=Promise.resolve();let he=!1;function Ze(){he||(he=!0,Ce.then(tt))}function $e(){return Ze(),Ce}function et(r){ue.push(r)}const de=new Set;let oe=0;function tt(){do{for(;oe<ne.length;){const r=ne[oe];oe++,lt(r.$$)}for(ne.length=0,oe=0;ke.length;)ke.pop()();for(let r=0;r<ue.length;r+=1){const e=ue[r];de.has(e)||(de.add(e),e())}ue.length=0}while(ne.length);for(;Ve.length;)Ve.pop()();he=!1,de.clear()}function lt(r){if(r.fragment!==null){r.update(),Ye(r.before_update);const e=r.dirty;r.dirty=[-1],r.fragment&&r.fragment.p(r.ctx,e),r.after_update.forEach(et)}}const st={root:null,rootMargin:"0px",threshold:0,unobserveOnEnter:!1},ce=(r,e)=>new CustomEvent(r,{detail:e});function re(r,e={}){const{root:a,rootMargin:g,threshold:n,unobserveOnEnter:i}=Object.assign(Object.assign({},st),e);let h={x:void 0,y:void 0},b={vertical:void 0,horizontal:void 0};if(typeof IntersectionObserver<"u"&&r){const d=new IntersectionObserver((w,o)=>{w.forEach(p=>{h.y>p.boundingClientRect.y?b.vertical="up":b.vertical="down",h.x>p.boundingClientRect.x?b.horizontal="left":b.horizontal="right",h={y:p.boundingClientRect.y,x:p.boundingClientRect.x};const m={inView:p.isIntersecting,entry:p,scrollDirection:b,node:r,observer:o};r.dispatchEvent(ce("change",m)),p.isIntersecting?(r.dispatchEvent(ce("enter",m)),i&&o.unobserve(r)):r.dispatchEvent(ce("leave",m))})},{root:a,rootMargin:g,threshold:n});return $e().then(()=>{r.dispatchEvent(ce("init",{observer:d,node:r}))}),d.observe(r),{destroy(){d.unobserve(r)}}}}function De(r){let e,a,g,n,i,h,b,d,w,o,p,m,k,y;return{c(){e=u("div"),a=u("h1"),g=R("Reward the community, fuel innovation"),n=O(),i=u("p"),h=R(`Gittips is a bot for GitHub that enables owners of open-source projects to reward
					contributors with cryptocurrency`),b=O(),d=u("div"),w=u("a"),o=R("Docs"),p=O(),m=u("a"),k=R("Explore"),this.h()},l(v){e=f(v,"DIV",{class:!0});var x=_(e);a=f(x,"H1",{class:!0});var H=_(a);g=G(H,"Reward the community, fuel innovation"),H.forEach(l),n=P(x),i=f(x,"P",{class:!0});var D=_(i);h=G(D,`Gittips is a bot for GitHub that enables owners of open-source projects to reward
					contributors with cryptocurrency`),D.forEach(l),b=P(x),d=f(x,"DIV",{class:!0});var E=_(d);w=f(E,"A",{class:!0,target:!0,rel:!0,href:!0});var V=_(w);o=G(V,"Docs"),V.forEach(l),p=P(E),m=f(E,"A",{class:!0,href:!0});var M=_(m);k=G(M,"Explore"),M.forEach(l),E.forEach(l),x.forEach(l),this.h()},h(){s(a,"class","text-5xl md:text-7xl font-bold text-white"),s(i,"class","text-xl py-10 text-gray-400"),s(w,"class","btn btn-wide btn-primary btn-outline text-white font-bold rounded-full capitalize"),s(w,"target","_blank"),s(w,"rel","noreferrer"),s(w,"href","https://deed-labs.gitbook.io/gittips/"),s(m,"class","btn btn-wide btn-primary text-white font-bold rounded-full capitalize"),s(m,"href","/explore/bounties"),s(d,"class","flex flex-col md:flex-row items-center justify-center gap-4"),s(e,"class","max-w-4xl")},m(v,x){L(v,e,x),t(e,a),t(a,g),t(e,n),t(e,i),t(i,h),t(e,b),t(e,d),t(d,w),t(w,o),t(d,p),t(d,m),t(m,k)},i(v){y||q(()=>{y=J(e,Me,{y:200,duration:1500}),y.start()})},o:K,d(v){v&&l(e)}}}function He(r){let e,a,g,n,i,h,b,d,w;return{c(){e=u("div"),a=u("h2"),g=R("Our solution does not interfere with your "),n=u("span"),i=R("usual workflow"),h=O(),b=u("p"),d=R("Create tasks and send payments without leaving GitHub"),this.h()},l(o){e=f(o,"DIV",{});var p=_(e);a=f(p,"H2",{class:!0});var m=_(a);g=G(m,"Our solution does not interfere with your "),n=f(m,"SPAN",{class:!0});var k=_(n);i=G(k,"usual workflow"),k.forEach(l),m.forEach(l),h=P(p),b=f(p,"P",{class:!0});var y=_(b);d=G(y,"Create tasks and send payments without leaving GitHub"),y.forEach(l),p.forEach(l),this.h()},h(){s(n,"class","text-primary"),s(a,"class","text-4xl text-white font-bold"),s(b,"class","text-2xl text-gray-400 mt-5")},m(o,p){L(o,e,p),t(e,a),t(a,g),t(a,n),t(n,i),t(e,h),t(e,b),t(b,d)},i(o){w||q(()=>{w=J(e,me,{duration:1500}),w.start()})},o:K,d(o){o&&l(e)}}}function Oe(r){let e,a,g,n,i,h,b,d,w,o,p,m,k,y,v,x,H,D;return{c(){e=u("div"),a=u("h1"),g=R("Automated Bounty Creation"),n=O(),i=u("p"),h=R("We automate the process of creating bounties from issues on GitHub."),b=O(),d=u("p"),w=R("Just add "),o=u("span"),p=R("bounty"),m=R(" label to an issue."),y=O(),v=u("div"),x=u("img"),this.h()},l(E){e=f(E,"DIV",{class:!0});var V=_(e);a=f(V,"H1",{class:!0});var M=_(a);g=G(M,"Automated Bounty Creation"),M.forEach(l),n=P(V),i=f(V,"P",{class:!0});var X=_(i);h=G(X,"We automate the process of creating bounties from issues on GitHub."),X.forEach(l),b=P(V),d=f(V,"P",{class:!0});var F=_(d);w=G(F,"Just add "),o=f(F,"SPAN",{class:!0});var Y=_(o);p=G(Y,"bounty"),Y.forEach(l),m=G(F," label to an issue."),F.forEach(l),V.forEach(l),y=P(E),v=f(E,"DIV",{class:!0});var N=_(v);x=f(N,"IMG",{class:!0,src:!0,alt:!0}),N.forEach(l),this.h()},h(){s(a,"class","text-4xl font-bold"),s(i,"class","text-xl text-gray-400 mt-5"),s(o,"class","badge badge-primary badge-outline"),s(d,"class","text-xl text-gray-400"),s(e,"class","w-full md:w-1/2"),s(x,"class","card border border-primary image-full w-full md:w-5/6"),pe(x.src,H=qe)||s(x,"src",H),s(x,"alt","issue example"),s(v,"class","flex justify-center w-full md:w-1/2")},m(E,V){L(E,e,V),t(e,a),t(a,g),t(e,n),t(e,i),t(i,h),t(e,b),t(e,d),t(d,w),t(d,o),t(o,p),t(d,m),L(E,y,V),L(E,v,V),t(v,x)},p:K,i(E){k||q(()=>{k=J(e,r[2],{x:-200,duration:1e3}),k.start()}),D||q(()=>{D=J(v,r[2],{x:200,duration:1e3}),D.start()})},o:K,d(E){E&&l(e),E&&l(y),E&&l(v)}}}function Pe(r){let e,a,g,n,i,h,b,d,w,o,p,m,k,y;return{c(){e=u("div"),a=u("img"),i=O(),h=u("div"),b=u("h1"),d=R("Budget Control"),w=O(),o=u("p"),p=R("Each organization and user has its own balance reserved on the smart contract. "),m=u("br"),k=R(`
						So they can't accidentally spend more than planned.`),this.h()},l(v){e=f(v,"DIV",{class:!0});var x=_(e);a=f(x,"IMG",{class:!0,src:!0,alt:!0}),x.forEach(l),i=P(v),h=f(v,"DIV",{class:!0});var H=_(h);b=f(H,"H1",{class:!0});var D=_(b);d=G(D,"Budget Control"),D.forEach(l),w=P(H),o=f(H,"P",{class:!0});var E=_(o);p=G(E,"Each organization and user has its own balance reserved on the smart contract. "),m=f(E,"BR",{}),k=G(E,`
						So they can't accidentally spend more than planned.`),E.forEach(l),H.forEach(l),this.h()},h(){s(a,"class","card border border-primary w-full md:w-5/6"),pe(a.src,g=Je)||s(a,"src",g),s(a,"alt","issue example"),s(e,"class","w-full md:w-1/2"),s(b,"class","text-4xl font-bold"),s(o,"class","text-xl text-gray-400 mt-5"),s(h,"class","w-full md:w-1/2")},m(v,x){L(v,e,x),t(e,a),L(v,i,x),L(v,h,x),t(h,b),t(b,d),t(h,w),t(h,o),t(o,p),t(o,m),t(o,k)},p:K,i(v){n||q(()=>{n=J(e,r[2],{x:-200,duration:1e3}),n.start()}),y||q(()=>{y=J(h,r[2],{x:200,duration:1e3}),y.start()})},o:K,d(v){v&&l(e),v&&l(i),v&&l(h)}}}function Re(r){let e,a,g,n,i,h,b,d,w,o,p,m,k,y;return{c(){e=u("div"),a=u("h1"),g=R("Wallet Integration"),n=O(),i=u("p"),h=R(`Gittips integrates with blockchain, allowing contributors to receive their rewards
						directly in their wallets. `),b=u("br"),d=R(` This makes it easy for contributors to use rewards as they
						see fit.`),o=O(),p=u("div"),m=u("img"),this.h()},l(v){e=f(v,"DIV",{class:!0});var x=_(e);a=f(x,"H1",{class:!0});var H=_(a);g=G(H,"Wallet Integration"),H.forEach(l),n=P(x),i=f(x,"P",{class:!0});var D=_(i);h=G(D,`Gittips integrates with blockchain, allowing contributors to receive their rewards
						directly in their wallets. `),b=f(D,"BR",{}),d=G(D,` This makes it easy for contributors to use rewards as they
						see fit.`),D.forEach(l),x.forEach(l),o=P(v),p=f(v,"DIV",{class:!0});var E=_(p);m=f(E,"IMG",{class:!0,src:!0,alt:!0}),E.forEach(l),this.h()},h(){s(a,"class","text-4xl font-bold"),s(i,"class","text-xl text-gray-400 mt-5"),s(e,"class","w-full md:w-1/2"),s(m,"class","card border border-primary image-full w-full md:w-5/6"),pe(m.src,k=Ke)||s(m,"src",k),s(m,"alt","issue example"),s(p,"class","flex justify-center w-full md:w-1/2")},m(v,x){L(v,e,x),t(e,a),t(a,g),t(e,n),t(e,i),t(i,h),t(i,b),t(i,d),L(v,o,x),L(v,p,x),t(p,m)},p:K,i(v){w||q(()=>{w=J(e,r[2],{x:-200,duration:1e3}),w.start()}),y||q(()=>{y=J(p,r[2],{x:200,duration:1e3}),y.start()})},o:K,d(v){v&&l(e),v&&l(o),v&&l(p)}}}function Ge(r){let e,a,g,n,i,h,b;return{c(){e=u("div"),a=u("h1"),g=R("Empower your developer community"),n=O(),i=u("a"),h=R("Try it out"),this.h()},l(d){e=f(d,"DIV",{});var w=_(e);a=f(w,"H1",{class:!0});var o=_(a);g=G(o,"Empower your developer community"),o.forEach(l),n=P(w),i=f(w,"A",{class:!0,target:!0,rel:!0,href:!0});var p=_(i);h=G(p,"Try it out"),p.forEach(l),w.forEach(l),this.h()},h(){s(a,"class","text-white text-4xl"),s(i,"class","btn btn-wide btn-primary text-white font-bold rounded-full mt-12"),s(i,"target","_blank"),s(i,"rel","noreferrer"),s(i,"href","https://deed-labs.gitbook.io/gittips/")},m(d,w){L(d,e,w),t(e,a),t(a,g),t(e,n),t(e,i),t(i,h)},i(d){b||q(()=>{b=J(e,me,{duration:1500}),b.start()})},o:K,d(d){d&&l(e)}}}function Le(r){let e,a,g,n,i,h,b,d,w,o,p,m,k,y,v,x,H;return{c(){e=u("div"),a=u("h1"),g=R("Roadmap 2023"),n=O(),i=u("ul"),h=u("li"),b=u("p"),d=R("Mainnet Launch"),w=O(),o=u("li"),p=u("p"),m=R("NFT Rewards Support"),k=O(),y=u("li"),v=u("p"),x=R("Deeper GitHub integration"),this.h()},l(D){e=f(D,"DIV",{class:!0});var E=_(e);a=f(E,"H1",{class:!0});var V=_(a);g=G(V,"Roadmap 2023"),V.forEach(l),n=P(E),i=f(E,"UL",{class:!0});var M=_(i);h=f(M,"LI",{"data-content":!0,class:!0});var X=_(h);b=f(X,"P",{class:!0});var F=_(b);d=G(F,"Mainnet Launch"),F.forEach(l),X.forEach(l),w=P(M),o=f(M,"LI",{"data-content":!0,class:!0});var Y=_(o);p=f(Y,"P",{class:!0});var N=_(p);m=G(N,"NFT Rewards Support"),N.forEach(l),Y.forEach(l),k=P(M),y=f(M,"LI",{"data-content":!0,class:!0});var Z=_(y);v=f(Z,"P",{class:!0});var U=_(v);x=G(U,"Deeper GitHub integration"),U.forEach(l),Z.forEach(l),M.forEach(l),E.forEach(l),this.h()},h(){s(a,"class","text-white text-4xl"),s(b,"class","w-full text-center"),s(h,"data-content","Q1"),s(h,"class","step step-primary"),s(p,"class","w-full text-center"),s(o,"data-content","Q1"),s(o,"class","step"),s(v,"class","w-full text-center"),s(y,"data-content","Q2"),s(y,"class","step"),s(i,"class","steps steps-vertical md:steps-horizontal mt-12 text-white"),s(e,"class","card bg-gray-700 w-full md:w-2/3 p-12")},m(D,E){L(D,e,E),t(e,a),t(a,g),t(e,n),t(e,i),t(i,h),t(h,b),t(b,d),t(i,w),t(i,o),t(o,p),t(p,m),t(i,k),t(i,y),t(y,v),t(v,x)},i(D){H||q(()=>{H=J(e,me,{duration:1500}),H.start()})},o:K,d(D){D&&l(e)}}}function rt(r){let e,a,g,n,i,h,b,d,w,o,p,m,k,y,v,x,H,D,E,V,M,X,F,Y,N,Z,U,ae,Q,$,fe,ve;e=new Qe({props:{class:"bg-base-100"}});let A=r[0]&&De(),W=r[1][0]&&He(),S=r[1][1]&&Oe(r),j=r[1][2]&&Pe(r),z=r[1][3]&&Re(r),B=r[1][4]&&Ge(),T=r[1][5]&&Le();return{c(){Ae(e.$$.fragment),a=O(),g=u("div"),n=u("div"),A&&A.c(),i=O(),h=u("div"),W&&W.c(),b=O(),d=u("div"),w=u("div"),o=u("ul"),p=u("li"),k=O(),y=u("li"),x=O(),H=u("li"),E=O(),V=u("div"),M=u("div"),S&&S.c(),X=O(),F=u("div"),j&&j.c(),Y=O(),N=u("div"),z&&z.c(),Z=O(),U=u("div"),B&&B.c(),ae=O(),Q=u("div"),T&&T.c(),this.h()},l(c){We(e.$$.fragment,c),a=P(c),g=f(c,"DIV",{class:!0});var I=_(g);n=f(I,"DIV",{class:!0});var _e=_(n);A&&A.l(_e),_e.forEach(l),I.forEach(l),i=P(c),h=f(c,"DIV",{class:!0});var be=_(h);W&&W.l(be),be.forEach(l),b=P(c),d=f(c,"DIV",{class:!0});var ie=_(d);w=f(ie,"DIV",{class:!0});var we=_(w);o=f(we,"UL",{class:!0});var ee=_(o);p=f(ee,"LI",{"data-content":!0,class:!0}),_(p).forEach(l),k=P(ee),y=f(ee,"LI",{"data-content":!0,class:!0}),_(y).forEach(l),x=P(ee),H=f(ee,"LI",{"data-content":!0,class:!0}),_(H).forEach(l),ee.forEach(l),we.forEach(l),E=P(ie),V=f(ie,"DIV",{class:!0});var te=_(V);M=f(te,"DIV",{class:!0});var xe=_(M);S&&S.l(xe),xe.forEach(l),X=P(te),F=f(te,"DIV",{class:!0});var ge=_(F);j&&j.l(ge),ge.forEach(l),Y=P(te),N=f(te,"DIV",{class:!0});var Ee=_(N);z&&z.l(Ee),Ee.forEach(l),te.forEach(l),ie.forEach(l),Z=P(c),U=f(c,"DIV",{class:!0});var ye=_(U);B&&B.l(ye),ye.forEach(l),ae=P(c),Q=f(c,"DIV",{class:!0});var Ie=_(Q);T&&T.l(Ie),Ie.forEach(l),this.h()},h(){s(n,"class","hero-content text-center"),s(g,"class","hero min-h-screen"),s(h,"class","flex flex-col items-center text-center justify-center p-5 md:p-28 h-64"),s(p,"data-content","✨"),s(p,"class",m="step "+(r[1][1]?"step-primary":"")),s(y,"data-content","💰"),s(y,"class",v="step "+(r[1][2]?"step-primary":"")),s(H,"data-content","💎"),s(H,"class",D="step "+(r[1][3]?"step-primary":"")),s(o,"class","steps steps-vertical h-full"),s(w,"class","flex flex-col justify-center items-center w-2/12 md:w-1/12"),s(M,"class","flex flex-col md:flex-row items-center gap-5 md:p-12 md:h-96"),s(F,"class","flex flex-col-reverse md:flex-row items-center gap-5 md:p-12 md:h-96"),s(N,"class","flex flex-col md:flex-row items-center gap-5 md:p-12 h-full md:h-96"),s(V,"class","w-10/12 md:w-11/12 p-4 md:p-0 flex flex-col gap-12 md:gap-5"),s(d,"class","flex flex-row items-stretch py-12 w-full max-w-screen-2xl m-auto text-white"),s(U,"class","py-24 text-center"),s(Q,"class","flex items-center justify-center max-w-screen-2xl w-full md:h-96 m-auto p-5 py-12")},m(c,I){Be(e,c,I),L(c,a,I),L(c,g,I),t(g,n),A&&A.m(n,null),L(c,i,I),L(c,h,I),W&&W.m(h,null),L(c,b,I),L(c,d,I),t(d,w),t(w,o),t(o,p),t(o,k),t(o,y),t(o,x),t(o,H),t(d,E),t(d,V),t(V,M),S&&S.m(M,null),t(V,X),t(V,F),j&&j.m(F,null),t(V,Y),t(V,N),z&&z.m(N,null),L(c,Z,I),L(c,U,I),B&&B.m(U,null),L(c,ae,I),L(c,Q,I),T&&T.m(Q,null),$=!0,fe||(ve=[le(re.call(null,h,{unobserveOnEnter:!0,rootMargin:"-20%"})),se(h,"change",r[3]),le(re.call(null,M,{unobserveOnEnter:!0,rootMargin:"-20%"})),se(M,"change",r[4]),le(re.call(null,F,{unobserveOnEnter:!0,rootMargin:"-20%"})),se(F,"change",r[5]),le(re.call(null,N,{unobserveOnEnter:!0,rootMargin:"-20%"})),se(N,"change",r[6]),le(re.call(null,U,{unobserveOnEnter:!0,rootMargin:"-20%"})),se(U,"change",r[7]),le(re.call(null,Q,{unobserveOnEnter:!0,rootMargin:"-20%"})),se(Q,"change",r[8])],fe=!0)},p(c,[I]){c[0]?A?I&1&&C(A,1):(A=De(),A.c(),C(A,1),A.m(n,null)):A&&(A.d(1),A=null),c[1][0]?W?I&2&&C(W,1):(W=He(),W.c(),C(W,1),W.m(h,null)):W&&(W.d(1),W=null),(!$||I&2&&m!==(m="step "+(c[1][1]?"step-primary":"")))&&s(p,"class",m),(!$||I&2&&v!==(v="step "+(c[1][2]?"step-primary":"")))&&s(y,"class",v),(!$||I&2&&D!==(D="step "+(c[1][3]?"step-primary":"")))&&s(H,"class",D),c[1][1]?S?(S.p(c,I),I&2&&C(S,1)):(S=Oe(c),S.c(),C(S,1),S.m(M,null)):S&&(S.d(1),S=null),c[1][2]?j?(j.p(c,I),I&2&&C(j,1)):(j=Pe(c),j.c(),C(j,1),j.m(F,null)):j&&(j.d(1),j=null),c[1][3]?z?(z.p(c,I),I&2&&C(z,1)):(z=Re(c),z.c(),C(z,1),z.m(N,null)):z&&(z.d(1),z=null),c[1][4]?B?I&2&&C(B,1):(B=Ge(),B.c(),C(B,1),B.m(U,null)):B&&(B.d(1),B=null),c[1][5]?T?I&2&&C(T,1):(T=Le(),T.c(),C(T,1),T.m(Q,null)):T&&(T.d(1),T=null)},i(c){$||(C(e.$$.fragment,c),C(A),C(W),C(S),C(j),C(z),C(B),C(T),$=!0)},o(c){Te(e.$$.fragment,c),$=!1},d(c){Fe(e,c),c&&l(a),c&&l(g),A&&A.d(),c&&l(i),c&&l(h),W&&W.d(),c&&l(b),c&&l(d),S&&S.d(),j&&j.d(),z&&z.d(),c&&l(Z),c&&l(U),B&&B.d(),c&&l(ae),c&&l(Q),T&&T.d(),fe=!1,Ne(ve)}}}function at(r,e,a){let g=!1;Ue(()=>a(0,g=!0));let n=new Array(6);return window.innerWidth<=500&&(n[1]=!0,n[2]=!0,n[3]=!0),[g,n,(m,k)=>window.innerWidth>500?Me(m,k):{},({detail:m})=>{a(1,n[0]=m.inView,n)},({detail:m})=>{window.innerWidth>500&&a(1,n[1]=m.inView,n)},({detail:m})=>{window.innerWidth>500&&a(1,n[2]=m.inView,n)},({detail:m})=>{window.innerWidth>500&&a(1,n[3]=m.inView,n)},({detail:m})=>{a(1,n[4]=m.inView,n)},({detail:m})=>{a(1,n[5]=m.inView,n)}]}class ot extends Se{constructor(e){super(),je(this,e,at,rt,ze,{})}}export{ot as default};
