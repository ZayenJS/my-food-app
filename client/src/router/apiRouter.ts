import { BaseRouter } from './baseRouter';

export enum ApiRoutes {
  root = 'root',
  recipes = 'recipes',
  rate_recipe = 'rate_recipe',

  foods = 'foods',
  delete_food = 'delete_food',

  brands = 'brands',
}

class ApiRouter extends BaseRouter<ApiRoutes> {
  private readonly url = import.meta.env.VITE_API_URL;

  public readonly routes: Record<ApiRoutes, { path: string }> = {
    root: { path: import.meta.env.VITE_BACKEND_URL },
    recipes: { path: `${this.url}/recipe` },
    rate_recipe: { path: `${this.url}/recipe/:id/rate` },
    foods: { path: `${this.url}/food` },
    delete_food: { path: `${this.url}/food/:id` },
    brands: { path: `${this.url}/brand` },
  };
}

export const apiRouter = new ApiRouter();
