import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
import { BrowserRouter } from "react-router";

const root = document.getElementById("root");
if (!root) {
	throw "#root not found!";
}
createRoot(root).render(
	<StrictMode>
		{/* React Router 宣言モードを使用します */}
		<BrowserRouter>
			<App />
		</BrowserRouter>
	</StrictMode>,
);
