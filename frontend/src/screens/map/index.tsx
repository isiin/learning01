import {
	Box,
	Button,
	FormControl,
	InputLabel,
	List,
	ListItem,
	ListItemButton,
	ListItemIcon,
	ListItemText,
	ListSubheader,
	MenuItem,
	Select,
	type SelectChangeEvent,
	Stack,
} from "@mui/material";
import { useCallback, useMemo, useState } from "react";
import MapView, {
	type FitBoundsControl,
	MapMarkersLayer,
	type MapTile,
	MapTileLayer,
	type MapViewOptions,
} from "../../components/MapView";
import "./index.css";
import type { Feature, FeatureCollection } from "geojson";
import { STORAGE_KEYS } from "../../constants/storageKeys";
import { useLocalStorage } from "../../hooks/useLocalStorage";
import type { ApiData, Customer, Surveyor, WorkZone } from "../../types/types";
import { fetchMockApiData } from "./demo";

// 地図の初期位置
const mapViewOptions: MapViewOptions = {
	initialPos: [43.06004, 141.3527],
	initialZoom: 10,
};
// 利用可能な地図タイル
type MapTileId =
	| "cyberjapandata_pale"
	| "cyberjapandata_std"
	| "cyberjapandata_photo"
	| "openstreetmap";

const mapTiles: Record<MapTileId, MapTile> = {
	cyberjapandata_pale: {
		name: "地理院タイル - 淡色地図",
		url: "https://cyberjapandata.gsi.go.jp/xyz/pale/{z}/{x}/{y}.png",
		maxZoom: 18,
		minZoom: 5,
		attribution: `&copy; <a href="https://maps.gsi.go.jp/development/ichiran.html">地理院タイル</a>`,
	},
	cyberjapandata_std: {
		name: "地理院タイル - 標準地図",
		url: "https://cyberjapandata.gsi.go.jp/xyz/std/{z}/{x}/{y}.png",
		maxZoom: 18,
		minZoom: 5,
		attribution: `&copy; <a href="https://maps.gsi.go.jp/development/ichiran.html">地理院タイル</a>`,
	},
	cyberjapandata_photo: {
		name: "地理院タイル - 写真",
		url: "https://cyberjapandata.gsi.go.jp/xyz/seamlessphoto/{z}/{x}/{y}.jpg",
		maxZoom: 18,
		minZoom: 2,
		attribution: `&copy; <a href="https://maps.gsi.go.jp/development/ichiran.html">地理院タイル</a>`,
	},
	openstreetmap: {
		name: "OpenStreetMap",
		url: "https://tile.openstreetmap.org/{z}/{x}/{y}.png",
		maxZoom: 19,
		minZoom: 2,
		attribution:
			"&copy; <a href='https://www.openstreetmap.org/copyright'>OpenStreetMapの貢献者</a>",
	},
} as const;

// TODO 視認性の高いカラーパレット (20色)
const COLOR_PALETTE = [
	"#e6194b", // Red
	"#3cb44b", // Green
	"#ffe119", // Yellow
	"#4363d8", // Blue
	"#f58231", // Orange
	"#911eb4", // Purple
	"#46f0f0", // Cyan
	"#f032e6", // Magenta
	"#bcf60c", // Lime
	"#fabebe", // Pink
	"#008080", // Teal
	"#e6beff", // Lavender
	"#9a6324", // Brown
	"#fffac8", // Beige
	"#800000", // Maroon
	"#aaffc3", // Mint
	"#808000", // Olive
	"#ffd8b1", // Apricot
	"#000075", // Navy
	"#808080", // Gray
];
const COLOR_UNASSIGNED = "#ffffff"; // 未割当
const COLOR_SELECTED = "#ff0000"; // 選択作業区
const COLOR_UNSELECTED = "#cccccc"; // 非選択作業区

// 地図選択
function MapTileSelect(props: {
	mapTileId: MapTileId;
	onChange: (mapTileId: MapTileId) => void;
}) {
	const handleChange = (event: SelectChangeEvent) => {
		const newValue = event.target.value as MapTileId;
		props.onChange(newValue);
	};

	const menuItems = Object.entries(mapTiles).map(([k, v]) => (
		<MenuItem key={k} value={k}>
			{v.name}
		</MenuItem>
	));

	return (
		<FormControl variant="outlined">
			<InputLabel id="map-tile-select-label">地図選択</InputLabel>
			<Select
				labelId="map-tile-select-label"
				id="map-tile-select"
				value={props.mapTileId}
				label="地図選択"
				onChange={handleChange}
			>
				{menuItems}
			</Select>
		</FormControl>
	);
}

// 色付きアイコンコンポーネント
function ColorIcon(props: { color: string }) {
	return (
		<Box
			sx={{
				height: "16px",
				width: "16px",
				borderRadius: 2,
				backgroundColor: props.color,
			}}
		/>
	);
}

// 調査員一覧
function SurveyorList(props: {
	surveyors: Surveyor[];
	colorMap: Map<string, string>;
}) {
	const menuItems = props.surveyors.map((item) => (
		<ListItem key={item.id}>
			<ListItemIcon sx={{ minWidth: 32 }}>
				<ColorIcon color={props.colorMap.get(item.id) || COLOR_UNASSIGNED} />
			</ListItemIcon>
			<ListItemText primary={item.name} />
		</ListItem>
	));

	return (
		<List dense subheader={<ListSubheader>調査員一覧</ListSubheader>}>
			{menuItems}
		</List>
	);
}

// 作業区一覧
function WorkZoneList(props: {
	workZones: WorkZone[];
	surveyors: Surveyor[];
	selectedWorkZoneId: string | null;
	onWorkZoneClick: (workZoneId: string) => void;
	onResetMapPosition: () => void;
}) {
	const surveyorMap = useMemo(
		() => new Map(props.surveyors.map((s) => [s.id, s.name])),
		[props.surveyors],
	);

	const handleListItemClick = (id: string) => {
		props.onWorkZoneClick(id);
	};

	const menuItems = props.workZones.map((item) => (
		<ListItemButton
			key={item.id}
			selected={props.selectedWorkZoneId === item.id}
			onClick={() => {
				handleListItemClick(item.id);
			}}
		>
			<ListItemText
				primary={item.name}
				secondary={`担当: ${surveyorMap.get(item.surveyorId) || "未割当"}`}
			/>
			{props.selectedWorkZoneId === item.id && (
				<ColorIcon color={COLOR_SELECTED} />
			)}
		</ListItemButton>
	));

	return (
		<>
			<List dense subheader={<ListSubheader>作業区一覧</ListSubheader>}>
				{menuItems}
			</List>
			<Button variant="text" onClick={props.onResetMapPosition}>
				位置リセット
			</Button>
		</>
	);
}

// 地図画面
export default function MapScreen() {
	// ロジックを vm (ViewModel) としてまとめて受け取る
	const vm = useMapScreenLogic();

	return (
		<Stack
			direction="column"
			spacing={1}
			sx={{
				height: "100%",
				width: "100%",
				padding: 2,
				boxSizing: "border-box",
			}}
		>
			<Stack direction="row" spacing={2} alignItems="center">
				<MapTileSelect
					mapTileId={vm.mapTileId}
					onChange={vm.handleMapTileChange}
				/>
				<Button variant="outlined" onClick={vm.handleDemo}>
					デモ
				</Button>
			</Stack>
			<Stack
				direction="row"
				spacing={0}
				sx={{ flexGrow: 1, overflow: "hidden" }}
			>
				<Box
					sx={{
						width: 240,
						overflowY: "auto",
						borderRight: "1px solid rgba(0, 0, 0, 0.12)",
					}}
				>
					<SurveyorList
						surveyors={vm.surveyors}
						colorMap={vm.surveyorColorMap}
					/>
				</Box>
				<Box
					sx={{
						width: 200,
						overflowY: "auto",
						borderRight: "1px solid rgba(0, 0, 0, 0.12)",
					}}
				>
					<WorkZoneList
						workZones={vm.workZones}
						surveyors={vm.surveyors}
						onWorkZoneClick={vm.handleWorkZoneClick}
						selectedWorkZoneId={vm.selectedWorkZoneId}
						onResetMapPosition={vm.handleResetMapPosition}
					/>
				</Box>
				<Box sx={{ flexGrow: 1, position: "relative" }}>
					<MapView options={mapViewOptions}>
						<MapTileLayer mapTile={mapTiles[vm.mapTileId]} />
						<MapMarkersLayer
							geoJson={vm.markerData}
							tooltipContent={getFeatureContent}
							popupContent={getFeatureContent}
							markerOptions={vm.getMarkerOptions}
							disableClusteringAtZoom={13}
							fitBoundsControl={vm.fitBoundsControl}
						/>
					</MapView>
				</Box>
			</Stack>
		</Stack>
	);
}

/**
 * MapScreenのロジック（ViewModel）
 */
const useMapScreenLogic = () => {
	const [mapTileId, setMapTileId] = useLocalStorage<MapTileId>(
		STORAGE_KEYS.MAP_TILE_ID,
		"cyberjapandata_pale",
	);
	const [apiData, setApiData] = useState<ApiData | null>(null);
	const [selectedWorkZoneId, setSelectedWorkZoneId] = useState<string | null>(
		null,
	);
	const [fitBoundsControl, setFitBoundsControl] = useState<FitBoundsControl>({
		trigger: 0,
	});

	// 地図選択
	const handleMapTileChange = (mapTileId: MapTileId) => {
		setMapTileId(mapTileId);
	};

	// 作業区Click
	const handleWorkZoneClick = (workZoneId: string) => {
		// すでに選択されている場合は解除(null)、そうでなければ選択
		setSelectedWorkZoneId((prev) => (prev === workZoneId ? null : workZoneId));
	};

	// 作業区DoubleClick
	const handleWorkZoneDoubleClick = (workZoneId: string) => {
		const targetPoints = apiData?.customers
			.filter((c) => c.workZoneId === workZoneId)
			.map<[number, number]>((c) => [c.lat, c.lng]);
		if (!targetPoints || targetPoints.length === 0) {
			return;
		}

		setFitBoundsControl((prev) => ({
			trigger: prev.trigger + 1,
			targetPoints: targetPoints,
		}));
	};

	// 位置リセット
	const handleResetMapPosition = () => {
		setFitBoundsControl((prev) => ({ trigger: prev.trigger + 1 }));
	};

	// デモ
	const handleDemo = async () => {
		try {
			// APIから一括でデータを取得
			const data = await fetchMockApiData();
			setApiData(data);
			setFitBoundsControl((prev) => ({ trigger: prev.trigger + 1 }));
		} catch (error) {
			console.error("Failed to fetch data:", error);
			// TODO: ユーザーにエラーを通知する処理（スナックバーなど）
		}
	};

	// 以下、useMemo, useCallbackを使用してメモ化。
	// メモ化を行わないとMapScreenが再レンダリングされるたびに再実行される。

	// 調査員リストから色マッピングを生成
	// 調査員リストが変更された場合のみ定義を更新する。
	const surveyorColorMap = useMemo(() => {
		if (!apiData) return new Map<string, string>();
		return createSurveyorColorMap(apiData.surveyors);
	}, [apiData]);

	// お客さまデータ(Raw)から表示用GeoJSONを生成
	// メモ化を行わないとMapScreenが再レンダリングされるたびにmarkerDataが再作成され
	// その結果MapMarkersLayerまで再レンダリングされる（propsが変更されるため）。
	const markerData = useMemo(() => {
		if (!apiData) return null;
		return convertToGeoJson(
			apiData.customers,
			apiData.surveyors,
			apiData.workZones,
		);
	}, [apiData]);

	// マーカーの表示形式
	// メモ化を行わないとMapScreenが再レンダリングされるたびにgetMarkerOptionsが再作成され
	// その結果MapMarkersLayerまで再レンダリングされる（propsが変更されるため）。
	const getMarkerOptions = useCallback(
		(feature: Feature) => {
			const id = feature.properties?.surveyorId;
			let color = surveyorColorMap.get(id) || COLOR_UNASSIGNED;
			if (selectedWorkZoneId) {
				if (feature.properties?.workZoneId === selectedWorkZoneId) {
					// 選択された作業区を強調
					color = COLOR_SELECTED;
				} else {
					// 選択された作業区以外をグレーアウト
					color = COLOR_UNSELECTED;
				}
			}
			return {
				type: "pin" as const,
				color: color,
			};
		},
		[surveyorColorMap, selectedWorkZoneId],
	);

	return {
		mapTileId,
		markerData,
		workZones: apiData?.workZones || [],
		surveyors: apiData?.surveyors || [],
		surveyorColorMap,
		selectedWorkZoneId,
		fitBoundsControl,
		handleMapTileChange,
		handleWorkZoneClick,
		handleWorkZoneDoubleClick,
		handleResetMapPosition,
		handleDemo,
		getMarkerOptions,
	};
};

// ツールチップ・ポップアップ用のコンテンツ生成関数
// コンポーネントの中に書くとメモ化する必要があるので外に出す
const getFeatureContent = (feature: Feature) => {
	const p = feature.properties;
	if (!p) return "";
	return `
		<b>ID:</b> ${p.id}<br />
		<b>お客さま:</b> ${p.name}<br />
		<b>担当:</b> ${p.surveyorName || "未割当"}<br />
		<b>作業区:</b> ${p.workZoneName || "不明"}
	`;
};

// お客さまデータをGeoJSONに変換するヘルパー関数（純粋関数）
const convertToGeoJson = (
	customers: Customer[],
	surveyors: Surveyor[],
	workZones: WorkZone[],
): FeatureCollection => {
	const surveyorMap = new Map(surveyors.map((s) => [s.id, s.name]));
	const workZoneMap = new Map(workZones.map((w) => [w.id, w.name]));

	return {
		type: "FeatureCollection",
		features: customers.map((item) => ({
			type: "Feature",
			geometry: { type: "Point", coordinates: [item.lng, item.lat] },
			properties: {
				...item,
				...{
					surveyorName: surveyorMap.get(item.surveyorId),
					workZoneName: workZoneMap.get(item.workZoneId),
				},
			},
		})),
	};
};

// 調査員リストから色マッピングを生成するヘルパー関数（純粋関数）
const createSurveyorColorMap = (surveyors: Surveyor[]) => {
	return new Map(
		surveyors.map((s, i) => [s.id, COLOR_PALETTE[i % COLOR_PALETTE.length]]),
	);
};
