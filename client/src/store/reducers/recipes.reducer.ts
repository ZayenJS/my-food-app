import { createReducer } from '@reduxjs/toolkit';

export type RecipesState = {
    list: unknown[];
};

const INITIAL_STATE: RecipesState = {
    list: [],
};

export const recipesReducer = createReducer(INITIAL_STATE, (builder) => builder);
