import { FC } from 'react';

export const NewFood: FC = () => {
    return (
        <div>
            <h2>Créer un aliment</h2>

            <form>
                <fieldset>
                    <legend>Informations générales</legend>
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
                    <div>
                        <label htmlFor="servings">Portions</label>
                        <input type="number" id="servings" name="servings" />
                    </div>
                    <div>
                        <label htmlFor="description">Description</label>
                        <textarea id="description" name="description" />
                    </div>
                    <div>
                        <label htmlFor="image">Image</label>
                        <input type="file" id="image" name="image" />
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
