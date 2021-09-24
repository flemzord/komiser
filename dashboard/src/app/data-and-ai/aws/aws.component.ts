import { Component, OnInit, OnDestroy } from '@angular/core';
import { AwsService } from '../../aws.service';
import { StoreService } from '../../store.service';
import { Subject, Subscription } from 'rxjs';
declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';


@Component({
  selector: 'aws-data-and-ai',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsDataAndAIComponent implements OnInit, OnDestroy {
  private sqsMessagesChart: any;

  public sqsQueues = 0;
  public numberOfMessagesSentToday = 0;
  public numberOfMessagesDeletedToday = 0;
  public snsTopics = 0;
  public activemqBrokers = 0;
  public kinesisStreams = 0;
  public kinesisShards = 0;
  public glueJobs = 0;
  public glueCrawlers = 0;
  public dataPipelines = 0;
  public esDomains = 0;
  public swfDomains = 0;

  private loadingSQS = true;
  private loadingSQSMessages = true;
  public loadingSNS = true;
  public loadingGlueCrawlers = true;
  public loadingActiveMQBrokers = true;
  public loadingGlueJobs = true;
  public loadingSwfDomains = true;
  public loadingDataPipelines = true;
  public loadingKinesisStreams = true;
  public loadingKinesisShards = true;
  public loadingESDomains = true;
  public loadingSQSMessagesChart = true;

  private _subscription: Subscription;

  constructor(private awsService: AwsService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(profile => {
      this.sqsMessagesChart.detach();

      const tooltips = document.getElementsByClassName('chartist-tooltip');
      for (let i = 0; i < tooltips.length; i++) {
        tooltips[i].outerHTML = '';
      }
      for (let j = 0; j < 3; j++) {
        const charts = document.getElementsByTagName('svg');
        for (let i = 0; i < charts.length; i++) {
          charts[i].outerHTML = '';
        }
      }

      this.sqsQueues = 0;
      this.numberOfMessagesSentToday = 0;
      this.numberOfMessagesDeletedToday = 0;
      this.snsTopics = 0;
      this.activemqBrokers = 0;
      this.kinesisStreams = 0;
      this.kinesisShards = 0;
      this.glueJobs = 0;
      this.glueCrawlers = 0;
      this.dataPipelines = 0;
      this.esDomains = 0;
      this.swfDomains = 0;

      this.loadingSQS = true;
      this.loadingSQSMessages = true;
      this.loadingSNS = true;
      this.loadingGlueCrawlers = true;
      this.loadingActiveMQBrokers = true;
      this.loadingGlueJobs = true;
      this.loadingSwfDomains = true;
      this.loadingDataPipelines = true;
      this.loadingKinesisStreams = true;
      this.loadingKinesisShards = true;
      this.loadingESDomains = true;
      this.loadingSQSMessagesChart = true;

      this.initState();
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  private initState() {
    this.awsService.getSQSPublishedMessagesMetrics().subscribe(data => {
      this.numberOfMessagesSentToday = data[0].Datapoints[Object.keys(data[0].Datapoints)[Object.keys(data[0].Datapoints).length - 1]];
      this.numberOfMessagesDeletedToday = data[1].Datapoints[Object.keys(data[1].Datapoints)[Object.keys(data[1].Datapoints).length - 1]];

      this.loadingSQSMessages = false;

      const seriesSent = [];
      const seriesDeleted = [];
      const labels = [];
      const i = 0;
      Object.keys(data[0].Datapoints).forEach(key => {
        labels.push(key);
        seriesSent.push({
          meta: 'NumberOfMessagesSent',
          value: data[0].Datapoints[key]
        });
        seriesDeleted.push({
          meta: 'NumberOfMessagesDeleted',
          value: data[1].Datapoints[key]
        });
      });


      this.loadingSQSMessagesChart = false;
      this.showSQSMessages(labels, [
        seriesSent,
        seriesDeleted
      ]);
    }, err => {
      this.loadingSQSMessagesChart = false;
      this.loadingSQSMessages = false;
    });

    this.awsService.getSQSQueues().subscribe(data => {
      this.sqsQueues = data;
      this.loadingSQS = false;
    }, err => {
      this.sqsQueues = 0;
      this.loadingSQS = false;
    });

    this.awsService.getSNSTopics().subscribe(data => {
      this.snsTopics = data;
      this.loadingSNS = false;
    }, err => {
      this.snsTopics = 0;
      this.loadingSNS = false;
    });

    this.awsService.getActiveMQBrokers().subscribe(data => {
      this.activemqBrokers = data;
      this.loadingActiveMQBrokers = false;
    }, err => {
      this.activemqBrokers = 0;
      this.loadingActiveMQBrokers = false;
    });

    this.awsService.getKinesisShards().subscribe(data => {
      this.kinesisShards = data;
      this.loadingKinesisShards = false;
    }, err => {
      this.kinesisShards = 0;
      this.loadingKinesisShards = false;
    });

    this.awsService.getKinesisStreams().subscribe(data => {
      this.kinesisStreams = data;
      this.loadingKinesisStreams = false;
    }, err => {
      this.kinesisStreams = 0;
      this.loadingKinesisStreams = false;
    });

    this.awsService.getGlueCrawlers().subscribe(data => {
      this.glueCrawlers = data;
      this.loadingGlueCrawlers = false;
    }, err => {
      this.glueCrawlers = 0;
      this.loadingGlueCrawlers = false;
    });

    this.awsService.getGlueJobs().subscribe(data => {
      this.glueJobs = data;
      this.loadingGlueJobs = false;
    }, err => {
      this.glueJobs = 0;
      this.loadingGlueJobs = false;
    });

    this.awsService.getDataPipelines().subscribe(data => {
      this.dataPipelines = data;
      this.loadingDataPipelines = false;
    }, err => {
      this.dataPipelines = 0;
      this.loadingDataPipelines = false;
    });

    this.awsService.getESDomains().subscribe(data => {
      this.esDomains = data;
      this.loadingESDomains = false;
    }, err => {
      this.esDomains = 0;
      this.loadingESDomains = false;
    });

    this.awsService.getSWFDomains().subscribe(data => {
      this.swfDomains = data;
      this.loadingSwfDomains = false;
    }, err => {
      this.swfDomains = 0;
      this.loadingSwfDomains = false;
    });
  }

  ngOnInit() {
  }

  private formatNumber(labelValue) {

    // Nine Zeroes for Billions
    return Math.abs(Number(labelValue)) >= 1.0e+9

      ? (Math.abs(Number(labelValue)) / 1.0e+9).toFixed(2) + ' B'
      // Six Zeroes for Millions
      : Math.abs(Number(labelValue)) >= 1.0e+6

        ? (Math.abs(Number(labelValue)) / 1.0e+6).toFixed(2) + ' M'
        // Three Zeroes for Thousands
        : Math.abs(Number(labelValue)) >= 1.0e+3

          ? (Math.abs(Number(labelValue)) / 1.0e+3).toFixed(2) + ' K'

          : Math.abs(Number(labelValue));
  }

  private showSQSMessages(labels, series) {
    const scope = this;
    this.sqsMessagesChart = new Chartist.Bar('#sqsMessagesChart', {
      labels: labels,
      series: series
    }, {
        plugins: [
          Chartist.plugins.tooltip()
        ],
        stackBars: true,
        axisY: {
          offset: 80,
          labelInterpolationFnc: function (value) {
            return scope.formatNumber(value);
          }
        }
      }).on('draw', function (data) {
        if (data.type === 'bar') {
          data.element.attr({
            style: 'stroke-width: 30px'
          });
        }
      });

  }
}
