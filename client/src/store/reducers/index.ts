import { combineReducers } from '@reduxjs/toolkit';
import { recipesReducer } from './recipes.reducer';
import { globalReducer } from './global.reducer';
import { brandReducer } from './brand.reducer';
import { foodReducer } from './food.reducer';

const rootReducer = combineReducers({
    global: globalReducer,
    brands: brandReducer,
    recipes: recipesReducer,
    foods: foodReducer,
});

export default rootReducer;
