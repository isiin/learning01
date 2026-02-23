import { Box, type SxProps, type Theme } from "@mui/material";
import {
	createContext,
	type ReactNode,
	useContext,
	useEffect,
	useRef,
	useState,
} from "react";
import "leaflet/dist/leaflet.css";
import "leaflet.markercluster/dist/MarkerCluster.css";
import "leaflet.markercluster/dist/MarkerCluster.Default.css";
import L from "leaflet";
import "leaflet.markercluster";
import type { Feature, FeatureCollection } from "geojson";
import pinSvg from "../assets/pin.svg?raw"; // SVGを文字列としてインポート

// MapViewのオプション
export type MapViewOptions = {
	initialPos: [number, number];
	initialZoom: number;
};

// 地図タイルの情報を格納する
export type MapTile = {
	name: string;
	url: string;
	minZoom?: number;
	maxZoom?: number;
	attribution: string;
};

// マーカーのオプション型定義
// Discriminated Union (判別可能な共用体) を使用して型安全にする
export type PinMarkerOptions = {
	type: "pin";
	color?: string;
};

export type CircleMarkerOptions = {
	type: "circle";
} & L.CircleMarkerOptions;

export type MarkerOptions = PinMarkerOptions | CircleMarkerOptions;

// Mapインスタンスを子コンポーネントに渡すためのContext
const MapContext = createContext<L.Map | null>(null);

// タイルレイヤーコンポーネント
export function MapTileLayer(props: { mapTile: MapTile }) {
	const map = useContext(MapContext);
	const layerRef = useRef<L.TileLayer | null>(null);

	useEffect(() => {
		if (!map) return;

		// レイヤーの差し替え
		if (layerRef.current) {
			map.removeLayer(layerRef.current);
		}
		layerRef.current = L.tileLayer(props.mapTile.url, {
			minZoom: props.mapTile.minZoom,
			maxZoom: props.mapTile.maxZoom,
			attribution: props.mapTile.attribution,
			crossOrigin: true,
		}).addTo(map);

		return () => {
			if (map && layerRef.current) {
				map.removeLayer(layerRef.current);
				layerRef.current = null;
			}
		};
	}, [map, props.mapTile]);

	return null;
}

// 色付きのピンアイコンを生成するヘルパー関数
const createPinIcon = (color: string) => {
	// SVG文字列内の黒色(#000000)を指定した色に置換
	const svg = pinSvg.replace(/#000000/g, color);

	return L.divIcon({
		className: "marker-pin-icon", // 既存のスタイル(枠線等)を適用させないためのダミークラス
		html: svg,
		iconSize: [36, 36],
		iconAnchor: [18, 36], // ピンの先端を座標に合わせる
		popupAnchor: [0, -36], // ポップアップの位置調整
	});
};

// ズーム制御用の型定義
export type FitBoundsControl = {
	trigger: number;
	targetPoints?: [number, number][]; // 指定がない場合はデータ全体の範囲とする
};

// マーカーレイヤーコンポーネント
export function MapMarkersLayer(props: {
	geoJson: FeatureCollection | null;
	tooltipContent?: (feature: Feature) => string;
	popupContent?: (feature: Feature) => string;
	markerOptions?: (feature: Feature) => MarkerOptions;
	disableClusteringAtZoom?: number;
	fitBoundsControl?: FitBoundsControl;
}) {
	const map = useContext(MapContext);
	const clusterRef = useRef<L.MarkerClusterGroup | null>(null);
	const geoJsonLayerRef = useRef<L.GeoJSON | null>(null);

	useEffect(() => {
		if (!map) return;

		// クラスタグループの初期化（初回のみ）
		if (!clusterRef.current) {
			clusterRef.current = L.markerClusterGroup({
				disableClusteringAtZoom: props.disableClusteringAtZoom,
			});
			map.addLayer(clusterRef.current);
		}

		const clusterGroup = clusterRef.current;
		clusterGroup.clearLayers();

		if (!props.geoJson) {
			geoJsonLayerRef.current = null;
			return;
		}

		const geoJsonLayer = L.geoJSON(props.geoJson, {
			pointToLayer: (feature, latlng) => {
				if (props.markerOptions) {
					const options = props.markerOptions(feature);
					if (options.type === "pin") {
						// 指定された色でピンアイコンを表示
						return L.marker(latlng, {
							icon: createPinIcon(options.color || "#3388ff"),
						});
					}
					// 指定された設定で円を表示
					return L.circleMarker(latlng, options);
				}
				// 標準のマーカー
				return L.marker(latlng);
			},
			onEachFeature: (feature, layer) => {
				if (props.tooltipContent) {
					const toolTipContent = props.tooltipContent(feature);
					if (toolTipContent) {
						layer.bindTooltip(toolTipContent);
					}
				}
				if (props.popupContent) {
					const popupContent = props.popupContent(feature);
					if (popupContent) {
						layer.bindPopup(popupContent);
					}
				}
			},
		});
		clusterGroup.addLayer(geoJsonLayer);
		geoJsonLayerRef.current = geoJsonLayer;
	}, [
		map,
		props.geoJson,
		props.tooltipContent,
		props.popupContent,
		props.markerOptions,
		props.disableClusteringAtZoom,
	]);

	useEffect(() => {
		if (!map) return;
		if (!geoJsonLayerRef.current) return;

		const control = props.fitBoundsControl;
		// 初期値(0)や未定義の場合は実行しない
		if (!control || !control.trigger) return;

		// targetPointsがあればその範囲に、なければデータ全体の範囲にズーム
		if (control.targetPoints && control.targetPoints.length > 0) {
			const bounds = L.latLngBounds(control.targetPoints);
			if (bounds.isValid()) map.fitBounds(bounds);
		} else {
			const bounds = geoJsonLayerRef.current.getBounds();
			if (bounds.isValid()) map.fitBounds(bounds);
		}
	}, [map, props.fitBoundsControl]);

	useEffect(() => {
		return () => {
			// コンポーネント破棄時に地図から削除
			if (map && clusterRef.current) {
				map.removeLayer(clusterRef.current);
				clusterRef.current = null;
			}
		};
	}, [map]);

	return null;
}

// 地図コンテナコンポーネント
export default function MapView(props: {
	children?: ReactNode;
	options?: MapViewOptions;
	sx?: SxProps<Theme>; //sxを親から指定できるようにする
}) {
	const mapDomRef = useRef<HTMLDivElement | null>(null);
	const mapRef = useRef<L.Map | null>(null);

	// 子コンポーネントに渡すためにStateで管理する
	const [map, setMap] = useState<L.Map | null>(null);

	// 初期化時のみ実行したいため、依存配列の警告を抑制
	// biome-ignore lint/correctness/useExhaustiveDependencies: props.options is used only for initialization
	useEffect(() => {
		const dom = mapDomRef.current;
		if (!dom) return;

		mapRef.current = L.map(dom);
		if (props.options) {
			mapRef.current.setView(
				props.options.initialPos,
				props.options.initialZoom,
			);
		} else {
			// オプションがない場合のデフォルト座標（例：札幌）
			mapRef.current.setView([43.06004, 141.3527], 18);
		}
		setMap(mapRef.current);

		return () => {
			if (mapRef.current) {
				mapRef.current.remove();
				mapRef.current = null;
				setMap(null);
			}
		};
	}, []);

	return (
		<MapContext.Provider value={map}>
			{/* sxを親から指定できるようにする */}
			<Box ref={mapDomRef} sx={{ height: "100%", width: "100%", ...props.sx }}>
				{props.children}
			</Box>
		</MapContext.Provider>
	);
}

// // 参考として残す。作業区レイヤーコンポーネント
// export function WorkZoneLayer(props: {
// 	geoJson: FeatureCollection | null;
// 	onZoneClick?: (zoneId: string) => void;
// }) {
// 	const map = useContext(MapContext);
// 	const layerRef = useRef<L.GeoJSON | null>(null);

// 	useEffect(() => {
// 		if (!map) return;

// 		if (layerRef.current) {
// 			map.removeLayer(layerRef.current);
// 			layerRef.current = null;
// 		}

// 		if (props.geoJson) {
// 			const layer = L.geoJSON(props.geoJson, {
// 				style: {
// 					color: "#3388ff",
// 					weight: 2,
// 					fillOpacity: 0.2,
// 				},
// 				onEachFeature: (feature, layer) => {
// 					// クリックイベント
// 					layer.on("click", () => {
// 						if (props.onZoneClick && feature.properties?.id) {
// 							props.onZoneClick(feature.properties.id);
// 						}
// 					});
// 				},
// 			}).addTo(map);
// 			layerRef.current = layer;
// 		}

// 		return () => {
// 			if (map && layerRef.current) {
// 				map.removeLayer(layerRef.current);
// 			}
// 		};
// 	}, [map, props.geoJson, props.onZoneClick]);

// 	return null;
// }
