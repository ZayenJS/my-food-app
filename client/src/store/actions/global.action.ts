import { createAction } from '@reduxjs/toolkit';

export enum GlobalActions {
  SET_PAGE_LOADING = 'global/setPageLoading',
  SET_PAGE_SEARCH = 'global/setPageSearch',
  SET_OPENED_MORE_MENU = 'global/setOpenedMoreMenu',
}

export const setPageLoading = createAction<boolean>(GlobalActions.SET_PAGE_LOADING);
export const setPageSearch = createAction<string>(GlobalActions.SET_PAGE_SEARCH);
export const setOpenedMoreMenu = createAction<string | null>(GlobalActions.SET_OPENED_MORE_MENU);
