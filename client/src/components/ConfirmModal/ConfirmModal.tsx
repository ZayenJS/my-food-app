import { FC, useRef, useState } from 'react';
import classes from './ConfirmModal.module.scss';
import { buildClassName } from '../../util';
import { LoadingButton } from '../LoadingButton/LoadingButton';

interface Props {
  text: string;
  visible: boolean;
  onConfirm: () => Promise<boolean>;
  onCancel: () => void;
}

export const ConfirmModal: FC<Props> = ({ text, visible, onConfirm, onCancel }) => {
  const [loading, setLoading] = useState(false);

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [error, setError] = useState<string | null>(null);
  const loadingButtonRef = useRef<HTMLButtonElement>(null);

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

  const onCancelHandler = () => {
    setError(null);
    onCancel();
  };

  const onConfirmHandler = async () => {
    setError(null);
    changeLoadingButtonStatus(true);
    const success = await onConfirm();
    changeLoadingButtonStatus(false);

    if (!success) {
      // TODO: display toast with error message
      setError('Une erreur est survenue lors de la confirmation, veuillez réessayer plus tard.');
    }
  };

  return (
    <>
      <button
        type="button"
        className={buildClassName(classes.overlay, visible ? classes.visible : '', loading ? classes.loading : '')}
        onClick={onCancelHandler}
      />
      <dialog className={buildClassName(classes.modal, visible ? classes.visible : '', loading ? classes.loading : '')}>
        <h2>Confirmation</h2>
        <div className={classes.icon}>
          <svg
            className={buildClassName(loading ? classes.loading : '')}
            onClick={onCancelHandler}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
          >
            <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z" />
          </svg>
        </div>
        <hr />
        <p>{text}</p>
        <div className={classes.buttons}>
          <button
            className={buildClassName('button', classes.button, classes.cancel, loading ? classes.loading : '')}
            onClick={onCancelHandler}
            type="button"
            disabled={loading}
          >
            Annuler
          </button>
          <LoadingButton
            type="button"
            ref={loadingButtonRef}
            classNames={{
              root: buildClassName('button', classes.button, classes.confirm),
              loading: buildClassName('loading', classes.loading),
            }}
            loading={loading}
            onClick={onConfirmHandler}
          >
            Confirmer
          </LoadingButton>
        </div>
      </dialog>
    </>
  );
};
