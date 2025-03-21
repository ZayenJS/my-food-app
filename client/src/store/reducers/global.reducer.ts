import { createReducer } from '@reduxjs/toolkit';
import { setPageLoading, setPageSearch } from '../actions/global.action';

export type GlobalState = {
    pageLoading: boolean;
    pageSearch: string;
};

const INITIAL_STATE: GlobalState = {
    pageLoading: false,
    pageSearch: '',
};

export const globalReducer = createReducer(INITIAL_STATE, (builder) =>
    builder
        .addCase(setPageLoading, (state, action) => {
            state.pageLoading = action.payload;
        })
        .addCase(setPageSearch, (state, action) => {
            state.pageSearch = action.payload;
        }),
);
