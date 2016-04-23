import {Component} from 'angular2/core';
import {RouteParams} from 'angular2/router';
import {Song} from './song';
import {SongService} from './song_service';

@Component({
  selector: 'song-create',
  template: `
      <form class="form-horizontal">
        <div class="form-group">
          <label for="inputTitle" class="col-sm-1 control-label">Title</label>
          <div class="col-sm-4">
            <input type="text" [(ngModel)]="song.name" class="form-control" id="inputTitle" placeholder="Title">
          </div>
        </div>
        <div class="form-group">
          <div class="col-sm-offset-1 col-sm-4">
            <button type="submit" class="btn btn-primary btn-sm" (click)="createSong()">Create</button>
          </div>
        </div>
      </form>`
})
export class SongCreate implements OnInit {
  song: Song;

  constructor(
      private routeParams: RouteParams,
      private songService: SongService) {}

  ngOnInit() {
    this.song = <Song>{}
  }

  createSong() {
    this.songService.createSong(this.song)
      .subscribe(
        (song: Song) => {
          this.song = song;
          window.history.back();
        },
        err => console.error(err)
      );
  }

  goBack() {
    window.history.back();
  }
}