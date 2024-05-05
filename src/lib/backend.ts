import { PUBLIC_BACKEND_URL, PUBLIC_LOBBY_API } from '$env/static/public';
import { HttpError } from './result';
import type { SimpleLobby, UserId, UserProfile } from './types';

type Fetch = typeof fetch;
const defaultFetch = fetch;

class Backend {
	private url: string;

	constructor(url: string) {
		this.url = url;
	}

	private async call<T = unknown>(
		method: string,
		path: string,
		body: Record<string, unknown> = {},
		fetch: Fetch = defaultFetch
	): Promise<T> {
		const headers = {};
		const result = await fetch(this.url + '/' + path, {
			method,
			mode: 'cors',
			credentials: 'include',
			headers,
			body: body && method !== 'GET' ? JSON.stringify(body) : undefined
		});
		const json = (await result.json()) as T;
		if (!result.ok) throw new HttpError(result.status, `Received error code ${result.status} from backend: ${JSON.stringify(json)}`);
		return json;
	}

	private async post<T = unknown>(path: string, body: Record<string, unknown> = {}, fetch: Fetch = defaultFetch): Promise<T> {
		return await this.call<T>('POST', path, body, fetch);
	}

	private async get<T = unknown>(path: string, body: Record<string, unknown> = {}, fetch: Fetch = defaultFetch): Promise<T> {
		const query = new URLSearchParams(Object.entries(body).map(([key, value]) => [key, JSON.stringify(value)]));
		return await this.call<T>('GET', `${path}?${query}`, {}, fetch);
	}

	// API CALLS

	async getProfile(fetch: Fetch = defaultFetch) {
		return await this.get<UserProfile>('v1/user/profile', {}, fetch);
	}

	async getUser(username: string, fetch: Fetch = defaultFetch) {
		return this.get<{
			id: UserId;
			name: string;
			username: string;
			email: string;
			avatar: string;
			background_img: string;
			bio: string;
			created_at: string;
			updated_at: string;
		}>(`v1/user/${username}`, {}, fetch);
	}

	async getUsers(fetch: Fetch = defaultFetch) {
		return this.get<{
			name: string;
			username: string;
			avatar: string;
			background_img: string;
			bio: string;
			created_at: string;
		}[]>(`v1/user`, {}, fetch);
	}

	// TODO move
	async getLobbies(): Promise<SimpleLobby[]> {
		const res = await fetch(`${PUBLIC_LOBBY_API}/lobbies`);
		const json = (await res.json()) as SimpleLobby[];
		return json
	}
}

export default new Backend(PUBLIC_BACKEND_URL);
