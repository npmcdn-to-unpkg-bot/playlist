import {SONGS} from './mock_songs';
import {Song} from './song'
import {Injectable} from 'angular2/core';
import {Http, Headers, Response} from 'angular2/http';

// Songs result.
interface SongsQueryResult {
	results: Song[],
	prevOffset: string,
	nextOffset: string
}

@Injectable()
export class SongService {
  constructor(public http: Http) {}

  getSongs() {
  	console.log('GET song list')
    return this.http.get('http://localhost:8080/v1.0/song/list')
      .map((res: Response) => res.json())
      .map((qr: SongsQueryResult) => {
      	let songs: Song[] = [];
      	if (qr) {
          qr.results.forEach((song) => {
          	songs.push(song);
          });
        }
        return songs;
      });
  }

  getSongDetail(id: string) {
  	let url = 'http://localhost:8080/v1.0/song/show/' + id

    return this.http.get(url)
      .map((res: Response) => res.json())
      .map((song: Song) => song);

      //.catch(this.handleError);
  }

  createSong(song: Song) {
  	let url = 'http://localhost:8080/v1.0/song/create'
  	let headers = new Headers();
    headers.append('Content-Type', 'application/json')

    return this.http.post(url, JSON.stringify(song), {headers: headers})
      .map((res: Response) => res.json());

      //.catch(this.handleError);
  }

  updateSong(song: Song) {
  	let url = 'http://localhost:8080/v1.0/song/update/' + song.id
  	let headers = new Headers();
    headers.append('Content-Type', 'application/json')

    return this.http.post(url, JSON.stringify(song), {headers: headers})
      .map((res: Response) => res.json());

      //.catch(this.handleError);
  }

  deleteSong(id: string) {
  	let url = 'http://localhost:8080/v1.0/song/delete/' + id
    return this.http.post(url, '', {})
      .map((res: Response) => res.json());

      //.catch(this.handleError);
  }

  private handleError (error: Response) {
    // Handle error.
    console.error(error);
    return Observable.throw(error.json().error || 'Server error');
  }
}