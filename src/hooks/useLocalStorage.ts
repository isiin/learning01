import { type Dispatch, type SetStateAction, useState } from "react";

/**
 * LocalStorageをuseStateのように扱うためのカスタムフック
 * @param key LocalStorageのキー
 * @param initialValue 初期値
 */
export function useLocalStorage<T>(
	key: string,
	initialValue: T,
): [T, Dispatch<SetStateAction<T>>] {
	// 初期値の取得
	const [storedValue, setStoredValue] = useState<T>(() => {
		try {
			const item = window.localStorage.getItem(key);
			// JSON.parse() は any 型を返すため、型アサーションで T 型であることを明示します。
			// これにより、useState が型を any と誤って解釈するのを防ぎます。
			// また、item がない場合は初期値を返します。
			return item ? (JSON.parse(item) as T) : initialValue;
		} catch (error) {
			console.warn(`Error reading localStorage key "${key}":`, error);
			return initialValue;
		}
	});

	// 値の更新用関数
	const setValue: Dispatch<SetStateAction<T>> = (value) => {
		try {
			// useStateと同様に、関数が渡された場合は現在の値を引数に実行する
			const valueToStore =
				value instanceof Function
					? (value as (val: T) => T)(storedValue)
					: value;
			setStoredValue(valueToStore);
			window.localStorage.setItem(key, JSON.stringify(valueToStore));
		} catch (error) {
			console.warn(`Error setting localStorage key "${key}":`, error);
		}
	};

	return [storedValue, setValue];
}
