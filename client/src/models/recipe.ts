import { TimeStamps } from './timestamps';

interface IngredientModel {
    food_id: number;
    quantity: number;
    unit: string;
    name: string;
}

interface StepModel {
    step_id: number;
    text: string;
    order: number;
}

interface TagModel {
    tag_id: number;
    name: string;
    color: string;
}

interface MacrosModel {
    calories_per_100g: number;
    protein_per_100g: number;
    carbs_per_100g: number;
    sugar_per_100g: number;
    fat_per_100g: number;
    saturated_per_100g: number;
    fiber_per_100g: number;
    sodium_per_100g: number;
}

export interface RecipeModel extends TimeStamps {
    recipe_id: number;
    name: string;
    description: string;
    difficulty: number;
    prep_time: number;
    cook_time: number;
    rest_time: number;
    servings: number;
    rating: number;
    image_url: string;
    ingredients: IngredientModel[];
    steps: StepModel[];
    tags: TagModel[];
    macros: MacrosModel;
}
