import { FC } from 'react';
import { FoodsIcon } from '../components/Icon/FoodsIcon/FoodsIcon';
import { MenuIcon } from '../components/Icon/MenuIcon/MenuIcon';
import { RecipeIcon } from '../components/Icon/RecipeIcon/RecipeIcon';
import { ShoppingListIcon } from '../components/Icon/ShoppingListIcon/ShoppingListIcon';
import { Foods } from '../pages/Foods/Foods';
import { Menus } from '../pages/Menus/Menus';
import { Recipes } from '../pages/Recipes/Recipes';
import { ShoppingLists } from '../pages/ShoppingLists/ShoppingLists';
import { Random } from '../util/random';
import { BaseRouter } from './baseRouter';
import { NewRecipe } from '../pages/NewRecipe/NewRecipe';
import { NewFood } from '../pages/NewFood/NewFood';

export enum AppRoutes {
  home = 'home',
  recipes = 'recipes',
  recipe_new = 'recipe_new',
  recipe = 'recipe',
  favorite_recipes = 'favorite_recipes',
  menus = 'menus',
  shoppingLists = 'shoppingLists',
  foods = 'foods',
  food_new = 'food_new',
  settings = 'settings',
}

class ClientRouter extends BaseRouter<AppRoutes> {
  public readonly routes: Record<
    AppRoutes,
    {
      path: string;
      component?: FC;
    }
  > = {
    home: { path: '/', component: Recipes },
    recipes: { path: '/recettes', component: Recipes },
    recipe_new: { path: '/recettes/creation', component: NewRecipe },
    recipe: { path: '/recettes/:id' }, // TODO: Add component
    favorite_recipes: { path: '/recettes/favoris' }, // TODO: Add component
    menus: { path: '/menus', component: Menus }, // TODO: Add component
    shoppingLists: { path: '/liste-de-courses', component: ShoppingLists },
    foods: { path: '/aliments', component: Foods },
    food_new: { path: '/aliments/creation', component: NewFood }, // TODO: Add component
    settings: { path: '/parametres' }, // TODO: Add component
  };

  public get sidebarNavigationLinks() {
    return [
      {
        id: Random.buildRandomId(),
        name: 'Recettes',
        href: this.buildLink(AppRoutes.recipes),
        icon: <RecipeIcon />,
      },
      {
        id: Random.buildRandomId(),
        name: 'Menus',
        href: this.buildLink(AppRoutes.menus),
        icon: <MenuIcon />,
      },
      {
        id: Random.buildRandomId(),
        name: 'Liste de courses',
        href: this.buildLink(AppRoutes.shoppingLists),
        icon: <ShoppingListIcon />,
      },
      {
        id: Random.buildRandomId(),
        name: 'Aliments',
        href: this.buildLink(AppRoutes.foods),
        icon: <FoodsIcon className="foods_icon" />,
      },
    ];
  }
}

export const clientRouter = new ClientRouter();
