import { Component, OnDestroy } from '@angular/core';
import { AwsService } from './aws.service';
import { StoreService } from './store.service';
import { not } from '@angular/compiler/src/output/output_ast';
import { Subscription } from 'rxjs';
import { Subject } from 'rxjs/Subject';
import * as moment from 'moment';

declare var ga: Function;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnDestroy {

  public accountName = 'Username';
  public redAlarms: number;
  public profiles: Array<string> = [];
  public currentProfile: string;
  public notifications: Array<Object> = [];
  public _subscription: Subscription;
  public currentProvider: any;
  public availableProviders: Array<any> = [
    {
      label: 'Amazon Web Services',
      value: 'aws'
    },
  ];

  private _storeService: StoreService;

  private providers: Map<String, Object> = new Map<String, Object>();

  constructor(private awsService: AwsService, private storeService: StoreService) {

    this.providers['aws'] = {
      label: 'Amazon Web Services',
      value: 'aws',
      logo: 'https://cdn.komiser.io/images/aws.png'
    };

    //if (this.storeService.getProvider() == 'aws') {
      if (localStorage.getItem('profile')) {
        this.currentProfile = localStorage.getItem('profile');
      } else {
        this.currentProfile = 'default';
        localStorage.setItem('profile', this.currentProfile);
      }

      this.awsService.getProfiles().subscribe(profiles => {
        this.profiles = profiles;
        if (this.profiles.length > 0 && this.profiles.indexOf(this.currentProfile) == -1) {
          this.currentProfile = this.profiles[0];
          localStorage.setItem('profile', this.currentProfile);
        }
      }, err => {
        this.profiles = [];
      });
   // }


    this.currentProvider = this.providers[this.storeService.getProvider()];
    this.storeService.onProviderChanged(this.storeService.getProvider());

    this._storeService = storeService;

    this.getAccountName();

    this._subscription = this.storeService.newNotification.subscribe(notifications => {
      this.notifications = [];
      Object.keys(notifications).forEach(key => {
        this.notifications.push(notifications[key]);
      });
    });
  }

  private getAccountName() {
    if (this.currentProvider.value == 'aws') {
      this.awsService.getAccountName().subscribe(data => {
        this.accountName = data.username;
      }, err => {
        this.accountName = 'Username';
      });

      this.awsService.getCloudwatchAlarms().subscribe(data => {
        this.redAlarms = data.ALARM;
      }, err => {
        this.redAlarms = 0;
      });
    }
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  public calcMoment(timestamp) {
    return moment(timestamp).fromNow();
  }

  public onCloudProviderSelected(provider) {
    this.currentProvider = this.providers[provider];
    this._storeService.onProviderChanged(provider);
    this.getAccountName();
  }

  public onProfileSelected(profile) {
    this.currentProfile = profile;
    localStorage.setItem('profile', this.currentProfile);
    this._storeService.onProfileChanged(profile);
    this.getAccountName();
  }

}
