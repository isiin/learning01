import type { ApiData, Customer } from "../../types/types";

// 中心座標
const CENTER_POS: [number, number] = [141.352, 43.06];

// 登場する調査員のID
const surveyorIds = [
	"000001",
	"000002",
	"100002",
	"100005",
	"801003",
	"820001",
	"820002",
];

// 作業区のテンプレート
const workZoneTemplates = [
	{
		id: "WZ-001",
		name: "中央区エリアA",
		centerOffsetKm: [0, 0], // 中心
		radiusKm: 0.8,
		count: 45,
		surveyorId: "000001",
	},
	{
		id: "WZ-002",
		name: "中央区エリアB",
		centerOffsetKm: [1.5, 0.5], // 東へ1.5km, 北へ0.5km
		radiusKm: 0.6,
		count: 30,
		surveyorId: "000002",
	},
	{
		id: "WZ-003",
		name: "北区エリアA",
		centerOffsetKm: [-0.5, 2.0], // 西へ0.5km, 北へ2.0km
		radiusKm: 1.0,
		count: 60,
		surveyorId: "801003",
	},
	{
		id: "WZ-004",
		name: "豊平区エリアA",
		centerOffsetKm: [1.0, -1.5], // 東へ1.0km, 南へ1.5km
		radiusKm: 0.7,
		count: 40,
		surveyorId: "",
	},
	{
		id: "WZ-005",
		name: "〇〇エリア1",
		centerOffsetKm: [1.5, -1.0],
		radiusKm: 0.2,
		count: 100,
		surveyorId: "820001",
	},
	{
		id: "WZ-006",
		name: "〇〇エリア2",
		centerOffsetKm: [1.5, -1.0],
		radiusKm: 0.2,
		count: 100,
		surveyorId: "820002",
	},
];

// APIからデータを取得するモック関数
export async function fetchMockApiData(): Promise<ApiData> {
	// 1. 調査員リスト
	const surveyors = surveyorIds.map((id) => ({
		id: id,
		name: `調査員${id}`,
	}));

	// 2. 作業区リスト
	const workZones = workZoneTemplates.map((zone) => ({
		id: zone.id,
		name: zone.name,
		surveyorId: zone.surveyorId,
	}));

	// 3. お客さまリスト
	const customers = generateCustomerData();

	return {
		surveyors,
		workZones,
		customers,
	};
}

// 各作業区テンプレートに配置されたお客さまリストを生成する
function generateCustomerData(): Customer[] {
	const data: Customer[] = [];
	let customerIdCounter = 1;

	for (const zoneTemplate of workZoneTemplates) {
		// 作業区の中心座標を計算（簡易的に緯度経度へ変換）
		// 1km ≒ 緯度0.009度, 経度0.012度(札幌付近)
		const zoneCenter: [number, number] = [
			CENTER_POS[0] + zoneTemplate.centerOffsetKm[0] * 0.012,
			CENTER_POS[1] + zoneTemplate.centerOffsetKm[1] * 0.009,
		];

		for (let i = 0; i < zoneTemplate.count; i++) {
			const coord = generateRandomPoint(zoneCenter, zoneTemplate.radiusKm);
			const id = String(customerIdCounter++);

			data.push({
				id: id,
				lng: coord[0],
				lat: coord[1],
				name: `お客さま${id}`,
				surveyorId: zoneTemplate.surveyorId,
				workZoneId: zoneTemplate.id,
			});
		}
	}

	return data;
}

/**
 * 指定座標を中心に、指定半径(km)内でランダムな座標を生成する
 * 円形に均等分布させる
 */
function generateRandomPoint(
	center: [number, number],
	radiusKm: number,
): [number, number] {
	const r = radiusKm / 111.32; // 半径を度数に変換（概算）
	const u = Math.random();
	const v = Math.random();
	const w = r * Math.sqrt(u);
	const t = 2 * Math.PI * v;
	const x = w * Math.cos(t);
	const y = w * Math.sin(t);

	// 経度の補正（緯度による距離の違いを考慮）
	const newLng = center[0] + x / Math.cos(center[1] * (Math.PI / 180));
	const newLat = center[1] + y;

	return [newLng, newLat];
}
