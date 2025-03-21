import { FC, FormEvent, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { FoodsIcon } from '../../components/Icon/FoodsIcon/FoodsIcon';

import classes from './NewFood.module.scss';
import { buildClassName } from '../../util';
import { ApiRoutes, apiRouter } from '../../router/apiRouter';
import { AppRoutes, clientRouter } from '../../router/clientRouter';
import { LoadingButton } from '../../components/LoadingButton/LoadingButton';
import { useAppDispatch } from '../../store';
import { setPageLoading } from '../../store/actions/global.action';
import { useBrands } from '../../hooks/useBrands';
import { createFood } from '../../store/actions/food.action';

export const NewFood: FC = () => {
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [loading, setLoading] = useState(false);
  const [brands] = useBrands();
  const [showImgPreview, setShowImgPreview] = useState(false);
  const appDispatch = useAppDispatch();

  const navigate = useNavigate();
  const loadingButtonRef = useRef<HTMLButtonElement>(null);
  const imgPreview = useRef<HTMLImageElement>(null);
  const imageInputRef = useRef<HTMLInputElement>(null);

  const onSubmitHandler = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const { method } = event.currentTarget;
    const url = event.currentTarget.action;
    const formData = new FormData(event.currentTarget);

    changeLoadingButtonStatus(true);

    const response = await fetch(url, {
      method,
      body: formData,
    });

    changeLoadingButtonStatus(false);

    if (!response.ok) {
      const json = await response.json();

      if (json.errors) {
        setErrors(json.errors);
        return;
      }

      setErrors(json);

      return;
    }

    appDispatch(setPageLoading(true));
    appDispatch(createFood(await response.json()));
    navigate(clientRouter.buildLink(AppRoutes.foods));
  };

  const changeLoadingButtonStatus = (status: boolean) => {
    const buttonWidth = loadingButtonRef.current?.offsetWidth;
    const buttonHeight = loadingButtonRef.current?.offsetHeight;

    if (!loadingButtonRef.current) {
      setLoading(status);
      return;
    }

    if (status) {
      loadingButtonRef.current.style.width = `${buttonWidth}px`;
      loadingButtonRef.current.style.height = `${buttonHeight}px`;
    } else {
      loadingButtonRef.current.style.width = ``;
      loadingButtonRef.current.style.height = ``;
    }

    setLoading(status);
  };

  const onImagePicked = (event: FormEvent<HTMLInputElement>) => {
    const file = (event.target as HTMLInputElement).files?.[0];

    if (!file) {
      return;
    }

    const reader = new FileReader();

    reader.onload = () => {
      if (imgPreview.current) {
        imgPreview.current.src = reader.result as string;
        setShowImgPreview(true);
      }
    };

    reader.readAsDataURL(file);
  };

  const onImageRemove = () => {
    if (imgPreview.current) {
      imgPreview.current.src = '';
      setShowImgPreview(false);
    }

    if (imageInputRef.current) {
      imageInputRef.current.value = '';
    }
  };

  return (
    <div className={buildClassName(classes.new_food, 'padded-wrapper-1rem')}>
      <h2 className="page_title">
        <FoodsIcon className="icon-lg" />
        <span>Créer un aliment</span>
      </h2>

      <form action={apiRouter.buildLink(ApiRoutes.foods)} method="POST" onSubmit={onSubmitHandler}>
        <fieldset>
          <legend>Informations générales</legend>

          <div className={buildClassName(classes.field_container, classes.general_info_container)}>
            <div>
              <label htmlFor="name">Nom</label>
              <input type="text" id="name" name="name" />
              {errors.name && <p className="error">{errors.name}</p>}
            </div>
            <div>
              <label htmlFor="brand">Marque</label>
              <input type="text" id="brand" list="brands" name="brand" />
              <datalist id="brands">
                {brands.map((brand) => (
                  <option key={brand.brand_id} value={brand.name} />
                ))}
              </datalist>
              {errors.brand && <p className="error">{errors.brand}</p>}
            </div>
            <div>
              <p>Image</p>
              <label htmlFor="image">Sélectionner une image</label>
              <input
                ref={imageInputRef}
                type="file"
                id="image"
                name="image"
                accept="image/png, image/jpeg, image/jpg"
                onChange={onImagePicked}
              />
            </div>
          </div>
          <div className={classes.img_preview} style={{ display: showImgPreview ? 'block' : 'none' }}>
            <img ref={imgPreview} alt="" />
            <button type="button" onClick={onImageRemove}>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" aria-hidden="true" className="v-icon__svg">
                <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"></path>
              </svg>
            </button>
          </div>
        </fieldset>

        <fieldset>
          <legend>Informations nutritionnelles</legend>

          <div className={classes.field_container}>
            <div>
              <label htmlFor="calories">Calories</label>
              <input type="number" id="calories" name="calories" />
              {errors.calories && <p className="error">{errors.calories}</p>}
            </div>
            <div>
              <label htmlFor="protein">Protéines (g)</label>
              <input type="number" id="protein" name="protein" />
              {errors.protein && <p className="error">{errors.protein}</p>}
            </div>
            <div>
              <label htmlFor="carbs">Glucides (g)</label>
              <input type="number" id="carbs" name="carbs" />
              {errors.carbs && <p className="error">{errors.carbs}</p>}
            </div>
            <div>
              <label htmlFor="sugar">Sucre (g)</label>
              <input type="number" id="sugar" name="sugar" />
              {errors.sugar && <p className="error">{errors.sugar}</p>}
            </div>
            <div>
              <label htmlFor="fat">Lipides (g)</label>
              <input type="number" id="fat" name="fat" />
              {errors.fat && <p className="error">{errors.fat}</p>}
            </div>
            <div>
              <label htmlFor="saturated_fat">Lipides saturés (g)</label>
              <input type="number" id="saturated_fat" name="saturated_fat" />
              {errors.saturated_fat && <p className="error">{errors.saturated_fat}</p>}
            </div>
            <div>
              <label htmlFor="fiber">Fibres (g)</label>
              <input type="number" id="fiber" name="fiber" />
              {errors.fiber && <p className="error">{errors.fiber}</p>}
            </div>
            <div>
              <label htmlFor="sodium">Sel (g)</label>
              <input type="number" id="sodium" name="sodium" />
              {errors.sodium && <p className="error">{errors.sodium}</p>}
            </div>
          </div>
        </fieldset>

        <LoadingButton
          type="submit"
          ref={loadingButtonRef}
          classNames={{
            root: 'button',
            loading: 'loading',
          }}
          loading={loading}>
          Créer
        </LoadingButton>
      </form>
    </div>
  );
};
