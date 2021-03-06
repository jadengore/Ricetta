define(function(require, exports, module) {

    var marionette = require('marionette');
    var template = require('hbs!../templates/curator-view')
    var RecipeCuratedView = require('app/views/recipe-curated-view').RecipeCuratedView;
    var Recipe = require('app/models/recipe').Recipe;

    var CuratorView = marionette.CompositeView.extend({
        childView: RecipeCuratedView,
        childViewContainer: '#curator',
        template: template,

        ui: {
        },

        initialize: function(options) {
            this.collection = options.collection;
            this.session = options.session;
            this.collection.fetch({
                success: function(res) {
                    console.log(res);
                }
            });
        }

    });

    exports.CuratorView = CuratorView;
})
