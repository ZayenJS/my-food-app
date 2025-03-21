import { createReducer } from '@reduxjs/toolkit';
import { setOpenedMoreMenu, setPageLoading, setPageSearch } from '../actions/global.action';

export type GlobalState = {
  pageLoading: boolean;
  pageSearch: string;
  openedMoreMenu: string | null;
};

const INITIAL_STATE: GlobalState = {
  pageLoading: false,
  pageSearch: '',
  openedMoreMenu: null,
};

export const globalReducer = createReducer(INITIAL_STATE, (builder) =>
  builder
    .addCase(setPageLoading, (state, action) => {
      state.pageLoading = action.payload;
    })
    .addCase(setPageSearch, (state, action) => {
      state.pageSearch = action.payload;
    })
    .addCase(setOpenedMoreMenu, (state, action) => {
      state.openedMoreMenu = action.payload;
    }),
);
