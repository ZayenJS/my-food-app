import { useEffect, useState } from 'react';

import { SearchInput } from '../../components/Form/SearchInput/SearchInput';
import { Recipe } from './Recipe/Recipe';

import { RecipeModel } from '../../models/recipe';
import { ApiRoutes, apiRouter } from '../../router/apiRouter';

import recipePlaceholder from '../../assets/images/recipe-placeholder.webp';

import classes from './Recipes.module.scss';

export const Recipes = () => {
    const [recipes, setRecipes] = useState<RecipeModel[]>([]);

    useEffect(() => {
        if (recipes.length > 0) return;

        const url = apiRouter.buildLink(ApiRoutes.recipes, {
            limit: 20,
        });

        fetch(url)
            .then((response) => response.json())
            .then((data: RecipeModel[]) => setRecipes(data));
    }, [recipes]);
    return (
        <div className={classes.root}>
            <div className={[classes.search_container, 'padded-wrapper-1rem'].join(' ')}>
                <SearchInput placeholder="Rechercher" />
            </div>

            <hr />

            <div className="padded-wrapper-1rem">
                <h2>
                    <span></span>
                    <span>Recettes</span>
                </h2>
            </div>
            <div className={[classes.recipes_container, 'padded-wrapper-1rem'].join(' ')}>
                {recipes.map((recipe) => (
                    <Recipe
                        key={recipe.id}
                        data={{
                            id: recipe.id,
                            title: recipe.name,
                            description: recipe.description,
                            isFavorite: false,
                            rating: recipe.rating,
                            src: recipe.image_url || recipePlaceholder,
                        }}
                    />
                ))}
            </div>
        </div>
    );
};
