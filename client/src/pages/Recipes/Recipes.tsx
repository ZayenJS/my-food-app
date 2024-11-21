import { useEffect, useState } from 'react';

import { Recipe } from './Recipe/Recipe';

import { RecipeModel } from '../../models/recipe';
import { ApiRoutes, apiRouter } from '../../router/apiRouter';

import recipePlaceholder from '../../assets/images/recipe-placeholder.webp';

import classes from './Recipes.module.scss';
import { buildClassName } from '../../util';
import { RecipeIcon } from '../../components/Icon/RecipeIcon/RecipeIcon';
import { PageSearch } from '../../components/PageSearch/PageSearch';

export const Recipes = () => {
    const [recipes, setRecipes] = useState<RecipeModel[]>([]);
    const [fetched, setFetched] = useState(false);

    useEffect(() => {
        if (recipes.length > 0 || fetched) return;

        const url = apiRouter.buildLink(ApiRoutes.recipes, {
            query: { limit: '20' },
        });

        fetch(url)
            .then((response) => {
                setFetched(true);
                if (!response.ok) {
                    throw new Error(response.statusText);
                }

                return response.json();
            })
            .then((data: RecipeModel[]) => setRecipes(data))
            .catch(() => setRecipes([]));
    }, [recipes, fetched]);

    return (
        <div className={classes.recipes}>
            <PageSearch />
            <div className="padded-wrapper-1rem">
                <h2 className="page_title">
                    <RecipeIcon className="icon-lg" />
                    <span>Recettes</span>
                </h2>
            </div>
            <div className={buildClassName(classes.recipes_container, 'padded-wrapper-1rem')}>
                {recipes.map((recipe) => (
                    <Recipe
                        key={recipe.recipe_id}
                        data={{
                            id: recipe.recipe_id,
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
