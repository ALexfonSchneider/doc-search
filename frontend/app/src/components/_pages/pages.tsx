import { FC } from "react";
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { SearchPage } from "./search-page/search-page";


export const Pages: FC = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" Component={SearchPage}/>
            </Routes>
        </BrowserRouter>
    )
}