import { Letter } from '../@types';
import { TimeStamps } from './timestamps';

export interface FoodModel extends TimeStamps {
  food_id: number;
  name: string;
  brand_id: number;
  calories: number;
  protein: number;
  carbs: number;
  sugar: number;
  fat: number;
  saturated_fat: number;
  fiber: number;
  sodium: number;
  image_url: string;
}

export type FoodModelIndexedByLetter = {
  [letter in Letter]?: FoodModel[];
};
