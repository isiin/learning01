// APIから以下のデータを取得すると仮定する

// お客さま
export type Customer = {
	id: string;
	lat: number;
	lng: number;
	name: string;
	surveyorId: string;
	workZoneId: string;
};

// 調査員
export type Surveyor = {
	id: string;
	name: string;
};

// 作業区
export type WorkZone = {
	id: string;
	name: string;
	surveyorId: string;
};

// APIレスポンスの型定義
export type ApiData = {
	surveyors: Surveyor[];
	workZones: WorkZone[];
	customers: Customer[];
};
