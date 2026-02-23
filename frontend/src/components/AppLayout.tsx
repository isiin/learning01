import {
	AppBar,
	Box,
	Divider,
	Drawer,
	List,
	ListItem,
	ListItemButton,
	ListItemText,
	Toolbar,
	Typography,
} from "@mui/material";
import { Link, Outlet } from "react-router";

const drawerWidth = 240;

function AppMenu() {
	return (
		<Drawer
			sx={{
				width: drawerWidth,
				flexShrink: 0,
				// Drawer内部のPaperコンポーネント（白いパネル部分）の幅を固定するための指定
				"& .MuiDrawer-paper": {
					width: drawerWidth,
					boxSizing: "border-box",
				},
			}}
			variant="permanent"
			anchor="left"
		>
			<Toolbar />
			<Divider />
			<List>
				{[
					{ text: "計画地図", url: "/map" },
					{ text: "実績地図", url: "/map" },
				].map((v) => (
					<ListItem key={v.text} disablePadding>
						<ListItemButton component={Link} to={v.url}>
							<ListItemText primary={v.text} />
						</ListItemButton>
					</ListItem>
				))}
			</List>
			<Divider />
			<List>
				{[{ text: "HOME", url: "/" }].map((v) => (
					<ListItem key={v.text} disablePadding>
						<ListItemButton component={Link} to={v.url}>
							<ListItemText primary={v.text} />
						</ListItemButton>
					</ListItem>
				))}
			</List>
		</Drawer>
	);
}

export default function AppLayout() {
	return (
		<Box
			sx={{
				display: "flex",
				height: "100vh",
				width: "100%",
				boxSizing: "border-box",
			}}
		>
			<AppBar
				position="fixed"
				sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}
			>
				<Toolbar variant="regular">
					<Typography variant="h6" noWrap component="div">
						My App
					</Typography>
				</Toolbar>
			</AppBar>
			<AppMenu />
			{/* メインコンテンツエリア */}
			<Box
				component="main"
				sx={{
					flexGrow: 1,
					display: "flex", // 縦並びのFlexコンテナにする
					flexDirection: "column",
					height: "100%", // 親要素(100vh)に合わせる
					overflow: "hidden", // 画面全体のスクロールを禁止
				}}
			>
				<Toolbar />
				{/* 残りの領域をすべて埋める。中身が溢れたらスクロールできるようにautoにする */}
				<Box sx={{ flexGrow: 1, overflow: "auto", position: "relative" }}>
					<Outlet />
				</Box>
			</Box>
		</Box>
	);
}
