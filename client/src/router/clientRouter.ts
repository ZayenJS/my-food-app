import { Random } from '../util/random';
import { BaseRouter } from './baseRouter';

export enum AppRoutes {
    home = 'home',
    recipes = 'recipes',
    recipe = 'recipe',
    menus = 'menus',
    shoppingList = 'shoppingList',
    settings = 'settings',
}

class ClientRouter extends BaseRouter {
    public readonly routes: Record<AppRoutes, string> = {
        home: '/',
        recipes: '/recettes',
        recipe: '/recettes/:id',
        menus: '/menus',
        shoppingList: '/liste-de-courses',
        settings: '/parametres',
    };

    public get sidebarNavigationLinks() {
        return [
            { id: Random.buildRandomId(), name: 'Recettes', href: this.buildLink(AppRoutes.recipes) },
            { id: Random.buildRandomId(), name: 'Menus', href: this.buildLink(AppRoutes.menus) },
            { id: Random.buildRandomId(), name: 'Liste de courses', href: this.buildLink(AppRoutes.shoppingList) },
        ];
    }
}

export const clientRouter = new ClientRouter();
