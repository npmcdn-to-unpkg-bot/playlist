import {Component} from 'angular2/core';
import {RouteConfig, ROUTER_DIRECTIVES, ROUTER_PROVIDERS} from 'angular2/router';
import {HTTP_PROVIDERS} from 'angular2/http';
import {SongList} from './song_list';
import {SongService} from './song_service';
import {SongDetail} from './song_detail';
import {SongCreate} from './song_create';

//    <song-list [songs]="songs"></song-list>
//    <song-form (newSong)="addSong($event)"></song-form>`,

@Component({
  selector: 'tunes-app',
  template: `
    <h5 class="text-muted">{{title}}</h5>
    <ul class="nav nav-tabs">
      <li role="presentation" class="active"><a [routerLink]="['Songs']">Songs</a></li>
    </ul>
    <p><p>
    <router-outlet></router-outlet>`,
  directives: [ROUTER_DIRECTIVES],
  providers: [
    ROUTER_PROVIDERS,
    HTTP_PROVIDERS,
    SongService
  ]
})
@RouteConfig([
  {
    path: '/song/list',
    name: 'Songs',
    component: SongList,
    useAsDefault: true
  },
  {
    path: '/song/show/:id',
    name: 'SongDetail',
    component: SongDetail
  },
  {
    path: '/song/create',
    name: 'SongCreate',
    component: SongCreate
  }
])
export class TunesApp {
  title = 'TUNES';
}