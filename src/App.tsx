import { CssBaseline, createTheme, ThemeProvider } from "@mui/material";
import { Route, Routes } from "react-router";
import AppLayout from "./components/AppLayout";
import HomeScreen from "./screens/home";
import MapScreen from "./screens/map";

// テーマを作成（必要に応じてここでフォントや色をカスタマイズします）
const theme = createTheme();

export default function App() {
	return (
		<ThemeProvider theme={theme}>
			<CssBaseline />
			<Routes>
				<Route path="/" element={<AppLayout />}>
					<Route index element={<HomeScreen />} />
					<Route path="map" element={<MapScreen />} />
				</Route>
			</Routes>
		</ThemeProvider>
	);
}
