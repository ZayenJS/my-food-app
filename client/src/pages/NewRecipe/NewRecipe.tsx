import { FC } from 'react';

import classes from './NewRecipe.module.scss';
import { buildClassName } from '../../util';
import { RecipeIcon } from '../../components/Icon/RecipeIcon/RecipeIcon';

export const NewRecipe: FC = () => {
    // const [ingredients, setIngredients] = useState<IngredientModel[]>([]);

    return (
        <div className={buildClassName(classes.new_recipe, 'padded-wrapper-1rem')}>
            <h2 className="page_title">
                <RecipeIcon className="icon-lg" />
                <span>Créer une recette</span>
            </h2>

            <form>
                <fieldset>
                    <legend>Informations générales</legend>
                    <div className={classes.general_info_container}>
                        <div>
                            <label htmlFor="name">Nom</label>
                            <input type="text" id="name" name="name" />
                        </div>
                        <div>
                            <label htmlFor="category">Catégorie</label>
                            <select id="category" name="category">
                                <option value="breakfast">Petit-déjeuner</option>
                                <option value="lunch">Déjeuner</option>
                                <option value="dinner">Dîner</option>
                                <option value="dessert">Dessert</option>
                                <option value="snack">Collation</option>
                            </select>
                        </div>

                        <div>
                            <label htmlFor="difficulty">Difficulté</label>
                            <select id="difficulty" name="difficulty">
                                <option value="1">Facile</option>
                                <option value="2">Moyen</option>
                                <option value="3">Difficile</option>
                            </select>
                        </div>
                    </div>
                    <div className={buildClassName(classes.field_container, classes.time_container)}>
                        <div>
                            <label htmlFor="prepTime">Temps de préparation (min)</label>
                            <input type="number" id="prepTime" name="prepTime" />
                        </div>
                        <div>
                            <label htmlFor="cookTime">Temps de cuisson (min)</label>
                            <input type="number" id="cookTime" name="cookTime" />
                        </div>
                        <div>
                            <label htmlFor="restTime">Temps de repos (min)</label>
                            <input type="number" id="restTime" name="restTime" />
                        </div>
                    </div>
                    <div className={classes.field_container}>
                        <div>
                            <label htmlFor="servings">Portions</label>
                            <input type="number" id="servings" name="servings" />
                        </div>
                    </div>
                    <div className={classes.field_container}>
                        <div>
                            <label htmlFor="description">Description</label>
                            <textarea id="description" name="description" />
                        </div>
                    </div>
                    <div className={buildClassName(classes.field_container, classes.image_container)}>
                        <div>
                            <label htmlFor="image">Image</label>
                            <input type="file" id="image" name="image" />
                        </div>
                    </div>
                </fieldset>
                <fieldset>
                    <legend>Ingrédients</legend>
                    {/*
                      TODO: Add ingredients list (dynamic)
                      Will trigger some sort of popup to search for ingredients
                    */}
                </fieldset>
                <fieldset>
                    <legend>Étapes</legend>
                    {/*
                      TODO: Add steps list (dynamic)
                      Will add a new step to the list when clicking on a button
                    */}
                </fieldset>
                <fieldset>
                    <legend>Tags</legend>
                    {/*
                      TODO: Add a button to add a new tag
                      Will trigger some sort of popup to search for tags
                      Add them to a list and display them
                    */}
                </fieldset>
                <div>
                    <button className="button" type="submit">
                        Créer
                    </button>
                </div>
            </form>
        </div>
    );
};
