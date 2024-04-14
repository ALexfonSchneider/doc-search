import { FC } from "react";
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { SearchPage } from "./search-page/search-page";
import { MetricsProvider } from "../_providers/metrics/metrics.provider";


export const Pages: FC = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<MetricsProvider><SearchPage/></MetricsProvider>}/>
            </Routes>
        </BrowserRouter>
    )
}