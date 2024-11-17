import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import './assets/scss/main.scss';
import { Provider } from 'react-redux';
import store from './store/index.ts';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Layout } from './components/Layout/Layout.tsx';
import { clientRouter, AppRoutes } from './router/router.ts';

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Provider store={store}>
            <BrowserRouter>
                <Layout>
                    <Routes>
                        {Object.entries(clientRouter.routes).map(([key, value]) => (
                            <Route
                                key={key}
                                path={clientRouter.buildLink(key as AppRoutes)}
                                Component={value.component}
                            />
                        ))}
                    </Routes>
                </Layout>
            </BrowserRouter>
        </Provider>
    </StrictMode>,
);
