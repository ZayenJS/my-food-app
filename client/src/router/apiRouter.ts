import { BaseRouter } from './baseRouter';

export enum ApiRoutes {
    recipes = 'recipes',
    rate_recipe = 'rate_recipe',
}

class ApiRouter extends BaseRouter {
    private readonly url = import.meta.env.VITE_API_URL;

    public readonly routes = {
        recipes: `${this.url}/recipe`,
        rate_recipe: `${this.url}/recipe/:id/rate`,
    };
}

export const apiRouter = new ApiRouter();
