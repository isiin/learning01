import {
	Box,
	Button,
	Container,
	Stack,
	type SxProps,
	type Theme,
	Typography,
} from "@mui/material";
import { useState } from "react";

// ▼ この画面でしか使わないパーツをファイル内で定義（exportはしない）
function PageHeader() {
	return (
		<Box>
			<Typography variant="h4" color="text.primary">
				Vite + React + TypeScript + MUI + React Router
			</Typography>
		</Box>
	);
}

function CounterArea(props: {
	count: number;
	onClick: () => void;
	sx?: SxProps<Theme>; //sxを親から指定できるようにする
}) {
	return (
		<Stack direction="column" alignItems="center" sx={props.sx}>
			<Button variant="contained" onClick={props.onClick}>
				こんにちは世界 {props.count}
			</Button>
			<Typography variant="body1" color="text.secondary">
				Edit <code>src/screens/home/index.tsx</code> and save to test HMR
			</Typography>
		</Stack>
	);
}

// ▼ メインの画面コンポーネント
export default function HomeScreen() {
	const [count, setCount] = useState(0);

	const handleClick = () => {
		setCount((prev) => prev + 1);
	};

	return (
		<>
			{/* Containerで囲むことで、中央寄せや最大幅が設定されます */}
			<Container maxWidth="lg" sx={{ height: "100%" }}>
				<Stack sx={{ alignItems: "center", p: 2 }} spacing={2}>
					<PageHeader />
					<CounterArea
						count={count}
						onClick={handleClick}
						sx={{ height: "500px", flexGrow: 1 }}
					/>
				</Stack>
			</Container>
		</>
	);
}
