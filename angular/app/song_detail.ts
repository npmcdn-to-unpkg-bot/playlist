import {Component, Output, EventEmitter} from 'angular2/core';
import {RouteParams} from 'angular2/router';
import {Song} from './song';
import {SongService} from './song_service';

//    <form (ngSubmit)="addSong()">
//      <input type="text" [(ngModel)]="name" size="30"
//             placeholder="Enter song title">
//      <input class="btn-primary" type="submit" value="add">
//    </form>

@Component({
  selector: 'song-detail',
  template: `
    <div *ngIf="song">
      <form class="form-horizontal">
        <div class="form-group">
          <label for="inputTitle" class="col-sm-1 control-label">Title</label>
          <div class="col-sm-4">
            <input type="text" [(ngModel)]="song.name" class="form-control" id="inputTitle" placeholder="Title">
          </div>
        </div>
        <div class="form-group">
          <div class="col-sm-offset-1 col-sm-4">
            <button type="submit" class="btn btn-primary btn-sm" (click)="updateSong()">Update</button>
            <button type="submit" class="btn btn-danger btn-sm" (click)="deleteSong()">Delete</button>
          </div>
        </div>
      </form>
    </div>`
})
export class SongDetail implements OnInit {
  song: Song;

  constructor(
    private routeParams: RouteParams,
    private songService: SongService) {}

  ngOnInit() {
    this.getSong()
  }

  saveSong(song: Song) {
    this.song = song
  }

  getSong() {
    this.songService.getSongDetail(this.routeParams.get('id'))
      .subscribe(
        (song: Song) => this.saveSong(song),
        err => console.error(err)
      );
  }

  updateSong() {
    this.songService.updateSong(this.song)
      .subscribe(
        (song: Song) => {
          this.saveSong(song);
          window.history.back();
        },
        err => console.error(err)
      );
  }

  deleteSong() {
    this.songService.deleteSong(this.routeParams.get('id'))
      .subscribe(
        () => {
          window.history.back();
        },
        err => console.error(err)
      );
  }

  goBack() {
    window.history.back();
  }
}