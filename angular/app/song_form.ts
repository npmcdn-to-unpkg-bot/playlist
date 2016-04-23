import {Component, Output, EventEmitter} from 'angular2/core';
import {Song} from './song';

@Component({
  selector: 'song-form',
  template: `
    <form (ngSubmit)="addSong()">
      <input type="text" [(ngModel)]="name" size="30"
             placeholder="Enter song title">
      <input class="btn-primary" type="submit" value="add">
    </form>`
})
export class SongForm {
  @Output() newSong = new EventEmitter<Song>();

  name: string = '';

  addSong() {
    if (this.name) {
      this.newSong.next({type: 'song', name: this.name, rank:0});
    }
    this.name = '';
  }
}