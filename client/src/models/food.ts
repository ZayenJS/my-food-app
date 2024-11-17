import { Letter } from '../@types';

export interface FoodModel {
    id: number;
    name: string;
    brandId: number;
    calories: number;
    protein: number;
    carbs: number;
    sugar: number;
    fat: number;
    saturatedFat: number;
    fiber: number;
    sodium: number;
    imageUrl: string;
    category: string;
    createdAt: string;
    updatedAt: string;
}

export type FoodModelIndexedByLetter = {
    [letter in Letter]?: FoodModel[];
};
