// @ts-nocheck
// This file is generated by Umi automatically
// DO NOT CHANGE IT MANUALLY!
import React from 'react';

export async function getRoutes() {
  const routes = {"1":{"path":"/user","layout":false,"id":"1"},"2":{"name":"login","path":"/user/login","parentId":"1","id":"2"},"3":{"path":"/welcome","name":"首页","icon":"smile","parentId":"ant-design-pro-layout","id":"3"},"4":{"path":"/","redirect":"/welcome","parentId":"ant-design-pro-layout","id":"4"},"5":{"path":"*","layout":false,"id":"5"},"6":{"path":"/user","name":"用户管理","icon":"crown","parentId":"ant-design-pro-layout","id":"6"},"7":{"path":"/user/userinfo","name":"用户信息","parentId":"6","id":"7"},"8":{"path":"/user/list","name":"用户列表","parentId":"6","id":"8"},"9":{"path":"/user/permission/resource/list","name":"资源列表","parentId":"6","id":"9"},"ant-design-pro-layout":{"id":"ant-design-pro-layout","path":"/","isLayout":true},"umi/plugin/openapi":{"path":"/umi/plugin/openapi","id":"umi/plugin/openapi"}} as const;
  return {
    routes,
    routeComponents: {
'1': React.lazy(() => import( './EmptyRoute')),
'2': React.lazy(() => import(/* webpackChunkName: "p__User__Login__index" */'@/pages/User/Login/index.tsx')),
'3': React.lazy(() => import(/* webpackChunkName: "p__Welcome" */'@/pages/Welcome.tsx')),
'4': React.lazy(() => import( './EmptyRoute')),
'5': React.lazy(() => import(/* webpackChunkName: "p__404" */'@/pages/404.tsx')),
'6': React.lazy(() => import( './EmptyRoute')),
'7': React.lazy(() => import(/* webpackChunkName: "p__User__UserInfo" */'@/pages/User/UserInfo.tsx')),
'8': React.lazy(() => import(/* webpackChunkName: "p__UserList__index" */'@/pages/UserList/index.tsx')),
'9': React.lazy(() => import(/* webpackChunkName: "p__PermissonResourceList__index" */'@/pages/PermissonResourceList/index.tsx')),
'ant-design-pro-layout': React.lazy(() => import(/* webpackChunkName: "umi__plugin-layout__Layout" */'/home/code/go-zero-antd-backend/web/src/.umi/plugin-layout/Layout.tsx')),
'umi/plugin/openapi': React.lazy(() => import(/* webpackChunkName: "umi__plugin-openapi__openapi" */'/home/code/go-zero-antd-backend/web/src/.umi/plugin-openapi/openapi.tsx')),
},
  };
}
