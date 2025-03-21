import { BaseRouter } from './baseRouter';

export enum ApiRoutes {
    recipes = 'recipes',
    rate_recipe = 'rate_recipe',

    foods = 'foods',

    brands = 'brands',
}

class ApiRouter extends BaseRouter<ApiRoutes> {
    private readonly url = import.meta.env.VITE_API_URL;

    public readonly routes: Record<ApiRoutes, { path: string }> = {
        recipes: { path: `${this.url}/recipe` },
        rate_recipe: { path: `${this.url}/recipe/:id/rate` },
        foods: { path: `${this.url}/food` },
        brands: { path: `${this.url}/brand` },
    };
}

export const apiRouter = new ApiRouter();
