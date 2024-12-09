import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const DASHBOARD: AppRouteRecordRaw = {
  path: '/layout',
  name: 'layout',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.dashboard',
    requiresAuth: true,
    icon: 'icon-dashboard',
    order: 0,
  },
  children: [
    {
      path: 'home',
      name: 'home',
      component: () => import('@/views/home/index.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'freAndTime',
      name: 'freAndTime',
      component: () => import('@/views/fre-time-chart/index.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'addNoise',
      name: 'addNoise',
      component: () => import('@/views/add-noise/index.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'filter',
      name: 'filter',
      component: () => import('@/views/filter/index.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'compare',
      name: 'compare',
      component: () => import('@/views/compare-chart/index.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default DASHBOARD;
