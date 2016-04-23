import {Component, OnInit, Pipe, PipeTransform} from 'angular2/core';
import {FORM_PROVIDERS, FORM_DIRECTIVES, Control} from 'angular2/common';
import {Router} from 'angular2/router';
import {RouteParams} from 'angular2/router';
import {Song} from './song';
import {SongService} from './song_service.ts';
import {SongDetail} from './song_detail';
import {SongCreate} from './song_create';

/*
    <div>
      <button type="submit" class="btn btn-default btn-sm" (click)="gotoCreate()">Create</button>
    </div>
    <div class="row">
          <div class="col-sm-1">#</div>
          <div class="col-sm-4">Title</div>
    </div>
*/

@Pipe({
  name: 'filter'
})
export class SongFilter implements PipeTransform {
  transform(songs: Song[], args: any[]) {
    let q: string = args[0].toLowerCase()
    return songs.filter((song) => song.name.toLowerCase().indexOf(q) >= 0)
  }
}

@Component({
  selector: 'song-list',

  pipes: [SongFilter],
  providers: [FORM_PROVIDERS],
  directives: [FORM_DIRECTIVES],

  template: `
    <div>
      <button type="submit" class="btn btn-info btn-sm" (click)="gotoCreate()">Create</button>
    </div>
    <p><p><p><p><p><p><p><p>
    <div class="row">
      <div class="col-sm-4">
        <input [ngFormControl]="filterTerm" class="form-control" (keyup)="applyFilter()" placeholder='Filter by name'/>
      </div>
    </div>
    <p><p><p><p><p><p><p><p>
    <div>
      <div class="row">
        <strong>
          <div class="col-sm-1">#</div>
          <div class="col-sm-4">Title</div>
        </strong>
      </div>
      <hr>
      <div *ngFor="#song of filteredSongs">
        <div class="row song" (click)="gotoDetail(song)">
          <div class="col-sm-1">{{song.id}}</div>
          <div class="col-sm-4">{{song.name}}</div>
        </div>
        <hr>
      </div>
    </div>`,

  styles:[`
    .song {
      background-color: white;
      color: black;
      display: block;
      height: 30px;
      line-height: 30px;
      text-decoration: none;
    }
    .song:hover {
      background-color: #f8f8f8;
    }
    hr{
      padding: 0px;
      margin: 0px;    
    }
  `]
})
export class SongList implements OnInit {
  songs: Song[] = [];
  filteredSongs: Song[] = [];
  filterTerm: Control = new Control();

  constructor(
    private router: Router,
    private routeParams: RouteParams,
    private songService: SongService) {}

  ngOnInit() {
    this.songService.getSongs()
      .subscribe(
        (songs: Song[]) => {
          this.songs = songs;
          this.filteredSongs = songs;
        },
        err => console.error(err)
      );
  }

  applyFilter() {
    let q: string = this.filterTerm.value.toLowerCase()
    this.filteredSongs = this.songs.filter((song) => song.name.toLowerCase().indexOf(q) >= 0)
  }

  gotoDetail(song: Song) {
    let link = ['SongDetail', {id: song.id}];
    this.router.navigate(link);
  }

  gotoCreate() {
    let link = ['SongCreate', {}];
    this.router.navigate(link);
  }
}

