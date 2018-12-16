import Main from '@/views/Main.vue';

// 不作为Main组件的子页面展示的页面单独写，如下
export const loginRouter = {
    path: '/login',
    name: 'login',
    meta: {
        title: 'Login - 登录'
    },
    component: () => import('@/views/login.vue')
};

// 作为Main组件的子页面展示但是不在左侧菜单显示的路由写在otherRouter里
export const otherRouter = {
    path: '/',
    name: 'otherRouter',
    redirect: '/home',
    component: Main,
    children: [
        {
            path: 'home',
            title: {i18n: 'home'},
            name: 'home_index',
            component: () => import('@/views/home/home.vue')
        },
        {
            path: 'ownspace',
            title: '个人中心',
            name: 'ownspace_index',
            component: () => import('@/views/own-space/own-space.vue')
        },
        {
            path: 'posts/add',
            title: '文章发布',
            name: 'add',
            component: () => import('@/views/posts/post-edit.vue')
        },
        {
            path: 'posts/edit/:id',
            title: '文章编辑',
            name: 'edit_post',
            component: () => import('@/views/posts/post-edit.vue')
        },
    ]
};

// 作为Main组件的子页面展示并且在左侧菜单显示的路由写在appRouter里
export const appRouter = [
    {
        path: '/dashboard',
        icon: 'home',
        name: 'dashboard',
        title: '首页',
        component: Main,
        children: [
            {
                path: 'home',
                icon: 'home',
                title: '首页',
                name: 'home',
                component: () => import('@/views/home/home.vue')
            }
        ]
    },
    {
        path: '/posts',
        icon: 'document-text',
        name: 'posts',
        title: '文章管理',
        component: Main,
        children: [
            {
                path: 'list',
                title: '文章管理',
                name: 'list',
                icon: 'document',
                component: () => import('@/views/posts/post-list.vue')
            },
            {
                path: 'categorys',
                title: '文章分类',
                name: 'categorys',
                icon: 'navicon',
                component: () => import('@/views/posts/categorys.vue')
            },
            {
                path: 'trash',
                title: '回收站',
                name: 'trash',
                icon: 'ios-trash',
                component: () => import('@/views/posts/post-trash.vue')
            },
            {
                path: 'tags',
                title: '文章标签',
                name: 'tags',
                icon: 'pricetags',
                component: () => import('@/views/posts/post-tags.vue')
            }
        ]
    },
    {
        path: '/extends',
        icon: 'hammer',
        name: 'extends',
        title: '扩展',
        component: Main,
        children: [
            {
                path: 'navigations',
                title: '导航管理',
                name: 'navigations',
                icon: 'ios-navigate-outline',
                component: () => import('@/views/plugins/navigations.vue')
            },
            {
                path: 'links',
                title: '左邻右舍',
                name: 'links',
                icon: 'link',
                component: () => import('@/views/plugins/links.vue')
            },
            {
                path: 'comments',
                title: '评论管理',
                name: 'comments',
                icon: 'chatboxes',
                component: () => import('@/views/plugins/comments.vue')
            },
            {
                path: 'options',
                title: '配置项目',
                name: 'options',
                icon: 'ios-settings',
                component: () => import('@/views/plugins/options.vue')
            }
        ]
    },
    {
        path: '/settings',
        icon: 'settings',
        name: 'settings',
        title: '系统设置',
        component: Main,
        children: [
            {
                path: 'setting',
                title: '系统设置',
                name: 'setting',
                component: () => import('@/views/plugins/settings.vue')
            }
        ]
    }
];


// 所有上面定义的路由都要写在下面的routers里
export const routers = [
    loginRouter,
    otherRouter,
    ...appRouter
];
