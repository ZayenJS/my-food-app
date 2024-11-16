import { combineReducers } from '@reduxjs/toolkit';
import { recipesReducer } from './recipes.reducer';

const rootReducer = combineReducers({ recipes: recipesReducer });

export default rootReducer;
