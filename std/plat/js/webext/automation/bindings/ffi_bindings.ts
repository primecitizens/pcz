import { importModule, Application, heap, Pointer } from "@ffi";

importModule("plat/js/webext/automation", (A: Application) => {
  const WEBEXT = typeof globalThis.browser === "undefined" ? globalThis.chrome : globalThis.browser;

  return {
    "store_Rect": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 20, false);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Bool(ptr + 18, false);
        A.store.Int32(ptr + 8, 0);
        A.store.Bool(ptr + 19, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 20, true);
        A.store.Bool(ptr + 16, "left" in x ? true : false);
        A.store.Int32(ptr + 0, x["left"] === undefined ? 0 : (x["left"] as number));
        A.store.Bool(ptr + 17, "top" in x ? true : false);
        A.store.Int32(ptr + 4, x["top"] === undefined ? 0 : (x["top"] as number));
        A.store.Bool(ptr + 18, "width" in x ? true : false);
        A.store.Int32(ptr + 8, x["width"] === undefined ? 0 : (x["width"] as number));
        A.store.Bool(ptr + 19, "height" in x ? true : false);
        A.store.Int32(ptr + 12, x["height"] === undefined ? 0 : (x["height"] as number));
      }
    },
    "load_Rect": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 16)) {
        x["left"] = A.load.Int32(ptr + 0);
      } else {
        delete x["left"];
      }
      if (A.load.Bool(ptr + 17)) {
        x["top"] = A.load.Int32(ptr + 4);
      } else {
        delete x["top"];
      }
      if (A.load.Bool(ptr + 18)) {
        x["width"] = A.load.Int32(ptr + 8);
      } else {
        delete x["width"];
      }
      if (A.load.Bool(ptr + 19)) {
        x["height"] = A.load.Int32(ptr + 12);
      } else {
        delete x["height"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_EventType": (ref: heap.Ref<string>): number => {
      const idx = [
        "accessKeyChanged",
        "activeDescendantChanged",
        "alert",
        "ariaAttributeChangedDeprecated",
        "ariaCurrentChanged",
        "atomicChanged",
        "autoCompleteChanged",
        "autocorrectionOccured",
        "autofillAvailabilityChanged",
        "blur",
        "busyChanged",
        "caretBoundsChanged",
        "checkedStateChanged",
        "checkedStateDescriptionChanged",
        "childrenChanged",
        "classNameChanged",
        "clicked",
        "collapsed",
        "controlsChanged",
        "detailsChanged",
        "describedByChanged",
        "descriptionChanged",
        "documentSelectionChanged",
        "documentTitleChanged",
        "dropeffectChanged",
        "editableTextChanged",
        "enabledChanged",
        "endOfTest",
        "expanded",
        "expandedChanged",
        "flowFromChanged",
        "flowToChanged",
        "focus",
        "focusAfterMenuClose",
        "focusChanged",
        "focusContext",
        "grabbedChanged",
        "haspopupChanged",
        "hide",
        "hierarchicalLevelChanged",
        "hitTestResult",
        "hover",
        "ignoredChanged",
        "imageAnnotationChanged",
        "imageFrameUpdated",
        "invalidStatusChanged",
        "keyShortcutsChanged",
        "labeledByChanged",
        "languageChanged",
        "layoutComplete",
        "layoutInvalidated",
        "liveRegionChanged",
        "liveRegionCreated",
        "liveRegionNodeChanged",
        "liveRelevantChanged",
        "liveStatusChanged",
        "loadComplete",
        "loadStart",
        "locationChanged",
        "mediaStartedPlaying",
        "mediaStoppedPlaying",
        "menuEnd",
        "menuItemSelected",
        "menuListValueChanged",
        "menuPopupEnd",
        "menuPopupStart",
        "menuStart",
        "mouseCanceled",
        "mouseDragged",
        "mouseMoved",
        "mousePressed",
        "mouseReleased",
        "multilineStateChanged",
        "multiselectableStateChanged",
        "nameChanged",
        "objectAttributeChanged",
        "orientationChanged",
        "otherAttributeChanged",
        "parentChanged",
        "placeholderChanged",
        "portalActivated",
        "positionInSetChanged",
        "rangeValueChanged",
        "rangeValueMaxChanged",
        "rangeValueMinChanged",
        "rangeValueStepChanged",
        "readonlyChanged",
        "relatedNodeChanged",
        "requiredStateChanged",
        "roleChanged",
        "rowCollapsed",
        "rowCountChanged",
        "rowExpanded",
        "scrollHorizontalPositionChanged",
        "scrollPositionChanged",
        "scrollVerticalPositionChanged",
        "scrolledToAnchor",
        "selectedChanged",
        "selectedChildrenChanged",
        "selectedValueChanged",
        "selection",
        "selectionAdd",
        "selectionRemove",
        "setSizeChanged",
        "show",
        "sortChanged",
        "stateChanged",
        "subtreeCreated",
        "textAttributeChanged",
        "textSelectionChanged",
        "textChanged",
        "tooltipClosed",
        "tooltipOpened",
        "treeChanged",
        "valueInTextFieldChanged",
        "valueChanged",
        "windowActivated",
        "windowDeactivated",
        "windowVisibilityChanged",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_ActionType": (ref: heap.Ref<string>): number => {
      const idx = [
        "annotatePageImages",
        "blur",
        "clearAccessibilityFocus",
        "collapse",
        "customAction",
        "decrement",
        "doDefault",
        "expand",
        "focus",
        "getImageData",
        "getTextLocation",
        "hideTooltip",
        "hitTest",
        "increment",
        "internalInvalidateTree",
        "loadInlineTextBoxes",
        "longClick",
        "replaceSelectedText",
        "resumeMedia",
        "scrollBackward",
        "scrollDown",
        "scrollForward",
        "scrollLeft",
        "scrollRight",
        "scrollUp",
        "scrollToMakeVisible",
        "scrollToPoint",
        "scrollToPositionAtRowColumn",
        "setAccessibilityFocus",
        "setScrollOffset",
        "setSelection",
        "setSequentialFocusNavigationStartingPoint",
        "setValue",
        "showContextMenu",
        "signalEndOfTest",
        "showTooltip",
        "startDuckingMedia",
        "stopDuckingMedia",
        "suspendMedia",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_IntentCommandType": (ref: heap.Ref<string>): number => {
      const idx = [
        "clearSelection",
        "delete",
        "dictate",
        "extendSelection",
        "format",
        "history",
        "insert",
        "marker",
        "moveSelection",
        "setSelection",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_IntentTextBoundaryType": (ref: heap.Ref<string>): number => {
      const idx = [
        "character",
        "formatEnd",
        "formatStart",
        "formatStartOrEnd",
        "lineEnd",
        "lineStart",
        "lineStartOrEnd",
        "object",
        "pageEnd",
        "pageStart",
        "pageStartOrEnd",
        "paragraphEnd",
        "paragraphStart",
        "paragraphStartSkippingEmptyParagraphs",
        "paragraphStartOrEnd",
        "sentenceEnd",
        "sentenceStart",
        "sentenceStartOrEnd",
        "webPage",
        "wordEnd",
        "wordStart",
        "wordStartOrEnd",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_IntentMoveDirectionType": (ref: heap.Ref<string>): number => {
      const idx = ["backward", "forward"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_AutomationIntent": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Enum(ptr + 4, -1);
        A.store.Enum(ptr + 8, -1);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(
          ptr + 0,
          [
            "clearSelection",
            "delete",
            "dictate",
            "extendSelection",
            "format",
            "history",
            "insert",
            "marker",
            "moveSelection",
            "setSelection",
          ].indexOf(x["command"] as string)
        );
        A.store.Enum(
          ptr + 4,
          [
            "character",
            "formatEnd",
            "formatStart",
            "formatStartOrEnd",
            "lineEnd",
            "lineStart",
            "lineStartOrEnd",
            "object",
            "pageEnd",
            "pageStart",
            "pageStartOrEnd",
            "paragraphEnd",
            "paragraphStart",
            "paragraphStartSkippingEmptyParagraphs",
            "paragraphStartOrEnd",
            "sentenceEnd",
            "sentenceStart",
            "sentenceStartOrEnd",
            "webPage",
            "wordEnd",
            "wordStart",
            "wordStartOrEnd",
          ].indexOf(x["textBoundary"] as string)
        );
        A.store.Enum(ptr + 8, ["backward", "forward"].indexOf(x["moveDirection"] as string));
      }
    },
    "load_AutomationIntent": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["command"] = A.load.Enum(ptr + 0, [
        "clearSelection",
        "delete",
        "dictate",
        "extendSelection",
        "format",
        "history",
        "insert",
        "marker",
        "moveSelection",
        "setSelection",
      ]);
      x["textBoundary"] = A.load.Enum(ptr + 4, [
        "character",
        "formatEnd",
        "formatStart",
        "formatStartOrEnd",
        "lineEnd",
        "lineStart",
        "lineStartOrEnd",
        "object",
        "pageEnd",
        "pageStart",
        "pageStartOrEnd",
        "paragraphEnd",
        "paragraphStart",
        "paragraphStartSkippingEmptyParagraphs",
        "paragraphStartOrEnd",
        "sentenceEnd",
        "sentenceStart",
        "sentenceStartOrEnd",
        "webPage",
        "wordEnd",
        "wordStart",
        "wordStartOrEnd",
      ]);
      x["moveDirection"] = A.load.Enum(ptr + 8, ["backward", "forward"]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_AutomationEvent_StopPropagation": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "stopPropagation" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationEvent_StopPropagation": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).stopPropagation);
    },

    "call_AutomationEvent_StopPropagation": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.stopPropagation();
    },
    "try_AutomationEvent_StopPropagation": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.stopPropagation();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_AutomationEvent_Target": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "target" in thiz) {
        const val = thiz.target;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationEvent_Target": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "target", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationEvent_Type": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "type" in thiz) {
        const val = thiz.type;
        A.store.Enum(
          retPtr,
          [
            "accessKeyChanged",
            "activeDescendantChanged",
            "alert",
            "ariaAttributeChangedDeprecated",
            "ariaCurrentChanged",
            "atomicChanged",
            "autoCompleteChanged",
            "autocorrectionOccured",
            "autofillAvailabilityChanged",
            "blur",
            "busyChanged",
            "caretBoundsChanged",
            "checkedStateChanged",
            "checkedStateDescriptionChanged",
            "childrenChanged",
            "classNameChanged",
            "clicked",
            "collapsed",
            "controlsChanged",
            "detailsChanged",
            "describedByChanged",
            "descriptionChanged",
            "documentSelectionChanged",
            "documentTitleChanged",
            "dropeffectChanged",
            "editableTextChanged",
            "enabledChanged",
            "endOfTest",
            "expanded",
            "expandedChanged",
            "flowFromChanged",
            "flowToChanged",
            "focus",
            "focusAfterMenuClose",
            "focusChanged",
            "focusContext",
            "grabbedChanged",
            "haspopupChanged",
            "hide",
            "hierarchicalLevelChanged",
            "hitTestResult",
            "hover",
            "ignoredChanged",
            "imageAnnotationChanged",
            "imageFrameUpdated",
            "invalidStatusChanged",
            "keyShortcutsChanged",
            "labeledByChanged",
            "languageChanged",
            "layoutComplete",
            "layoutInvalidated",
            "liveRegionChanged",
            "liveRegionCreated",
            "liveRegionNodeChanged",
            "liveRelevantChanged",
            "liveStatusChanged",
            "loadComplete",
            "loadStart",
            "locationChanged",
            "mediaStartedPlaying",
            "mediaStoppedPlaying",
            "menuEnd",
            "menuItemSelected",
            "menuListValueChanged",
            "menuPopupEnd",
            "menuPopupStart",
            "menuStart",
            "mouseCanceled",
            "mouseDragged",
            "mouseMoved",
            "mousePressed",
            "mouseReleased",
            "multilineStateChanged",
            "multiselectableStateChanged",
            "nameChanged",
            "objectAttributeChanged",
            "orientationChanged",
            "otherAttributeChanged",
            "parentChanged",
            "placeholderChanged",
            "portalActivated",
            "positionInSetChanged",
            "rangeValueChanged",
            "rangeValueMaxChanged",
            "rangeValueMinChanged",
            "rangeValueStepChanged",
            "readonlyChanged",
            "relatedNodeChanged",
            "requiredStateChanged",
            "roleChanged",
            "rowCollapsed",
            "rowCountChanged",
            "rowExpanded",
            "scrollHorizontalPositionChanged",
            "scrollPositionChanged",
            "scrollVerticalPositionChanged",
            "scrolledToAnchor",
            "selectedChanged",
            "selectedChildrenChanged",
            "selectedValueChanged",
            "selection",
            "selectionAdd",
            "selectionRemove",
            "setSizeChanged",
            "show",
            "sortChanged",
            "stateChanged",
            "subtreeCreated",
            "textAttributeChanged",
            "textSelectionChanged",
            "textChanged",
            "tooltipClosed",
            "tooltipOpened",
            "treeChanged",
            "valueInTextFieldChanged",
            "valueChanged",
            "windowActivated",
            "windowDeactivated",
            "windowVisibilityChanged",
          ].indexOf(val as string)
        );
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationEvent_Type": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(
        thiz,
        "type",
        val > 0 && val <= 119
          ? [
              "accessKeyChanged",
              "activeDescendantChanged",
              "alert",
              "ariaAttributeChangedDeprecated",
              "ariaCurrentChanged",
              "atomicChanged",
              "autoCompleteChanged",
              "autocorrectionOccured",
              "autofillAvailabilityChanged",
              "blur",
              "busyChanged",
              "caretBoundsChanged",
              "checkedStateChanged",
              "checkedStateDescriptionChanged",
              "childrenChanged",
              "classNameChanged",
              "clicked",
              "collapsed",
              "controlsChanged",
              "detailsChanged",
              "describedByChanged",
              "descriptionChanged",
              "documentSelectionChanged",
              "documentTitleChanged",
              "dropeffectChanged",
              "editableTextChanged",
              "enabledChanged",
              "endOfTest",
              "expanded",
              "expandedChanged",
              "flowFromChanged",
              "flowToChanged",
              "focus",
              "focusAfterMenuClose",
              "focusChanged",
              "focusContext",
              "grabbedChanged",
              "haspopupChanged",
              "hide",
              "hierarchicalLevelChanged",
              "hitTestResult",
              "hover",
              "ignoredChanged",
              "imageAnnotationChanged",
              "imageFrameUpdated",
              "invalidStatusChanged",
              "keyShortcutsChanged",
              "labeledByChanged",
              "languageChanged",
              "layoutComplete",
              "layoutInvalidated",
              "liveRegionChanged",
              "liveRegionCreated",
              "liveRegionNodeChanged",
              "liveRelevantChanged",
              "liveStatusChanged",
              "loadComplete",
              "loadStart",
              "locationChanged",
              "mediaStartedPlaying",
              "mediaStoppedPlaying",
              "menuEnd",
              "menuItemSelected",
              "menuListValueChanged",
              "menuPopupEnd",
              "menuPopupStart",
              "menuStart",
              "mouseCanceled",
              "mouseDragged",
              "mouseMoved",
              "mousePressed",
              "mouseReleased",
              "multilineStateChanged",
              "multiselectableStateChanged",
              "nameChanged",
              "objectAttributeChanged",
              "orientationChanged",
              "otherAttributeChanged",
              "parentChanged",
              "placeholderChanged",
              "portalActivated",
              "positionInSetChanged",
              "rangeValueChanged",
              "rangeValueMaxChanged",
              "rangeValueMinChanged",
              "rangeValueStepChanged",
              "readonlyChanged",
              "relatedNodeChanged",
              "requiredStateChanged",
              "roleChanged",
              "rowCollapsed",
              "rowCountChanged",
              "rowExpanded",
              "scrollHorizontalPositionChanged",
              "scrollPositionChanged",
              "scrollVerticalPositionChanged",
              "scrolledToAnchor",
              "selectedChanged",
              "selectedChildrenChanged",
              "selectedValueChanged",
              "selection",
              "selectionAdd",
              "selectionRemove",
              "setSizeChanged",
              "show",
              "sortChanged",
              "stateChanged",
              "subtreeCreated",
              "textAttributeChanged",
              "textSelectionChanged",
              "textChanged",
              "tooltipClosed",
              "tooltipOpened",
              "treeChanged",
              "valueInTextFieldChanged",
              "valueChanged",
              "windowActivated",
              "windowDeactivated",
              "windowVisibilityChanged",
            ][val - 1]
          : undefined,
        thiz
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationEvent_EventFrom": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "eventFrom" in thiz) {
        const val = thiz.eventFrom;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationEvent_EventFrom": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "eventFrom", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationEvent_MouseX": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "mouseX" in thiz) {
        const val = thiz.mouseX;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationEvent_MouseX": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "mouseX", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationEvent_MouseY": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "mouseY" in thiz) {
        const val = thiz.mouseY;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationEvent_MouseY": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "mouseY", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationEvent_Intents": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "intents" in thiz) {
        const val = thiz.intents;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationEvent_Intents": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "intents", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "constof_RoleType": (ref: heap.Ref<string>): number => {
      const idx = [
        "abbr",
        "alert",
        "alertDialog",
        "application",
        "article",
        "audio",
        "banner",
        "blockquote",
        "button",
        "canvas",
        "caption",
        "caret",
        "cell",
        "checkBox",
        "client",
        "code",
        "colorWell",
        "column",
        "columnHeader",
        "comboBoxGrouping",
        "comboBoxMenuButton",
        "comboBoxSelect",
        "comment",
        "complementary",
        "contentDeletion",
        "contentInsertion",
        "contentInfo",
        "date",
        "dateTime",
        "definition",
        "descriptionList",
        "descriptionListDetail",
        "descriptionListTerm",
        "desktop",
        "details",
        "dialog",
        "directory",
        "disclosureTriangle",
        "docAbstract",
        "docAcknowledgments",
        "docAfterword",
        "docAppendix",
        "docBackLink",
        "docBiblioEntry",
        "docBibliography",
        "docBiblioRef",
        "docChapter",
        "docColophon",
        "docConclusion",
        "docCover",
        "docCredit",
        "docCredits",
        "docDedication",
        "docEndnote",
        "docEndnotes",
        "docEpigraph",
        "docEpilogue",
        "docErrata",
        "docExample",
        "docFootnote",
        "docForeword",
        "docGlossary",
        "docGlossRef",
        "docIndex",
        "docIntroduction",
        "docNoteRef",
        "docNotice",
        "docPageBreak",
        "docPageFooter",
        "docPageHeader",
        "docPageList",
        "docPart",
        "docPreface",
        "docPrologue",
        "docPullquote",
        "docQna",
        "docSubtitle",
        "docTip",
        "docToc",
        "document",
        "embeddedObject",
        "emphasis",
        "feed",
        "figcaption",
        "figure",
        "footer",
        "footerAsNonLandmark",
        "form",
        "genericContainer",
        "graphicsDocument",
        "graphicsObject",
        "graphicsSymbol",
        "grid",
        "group",
        "header",
        "headerAsNonLandmark",
        "heading",
        "iframe",
        "iframePresentational",
        "image",
        "imeCandidate",
        "inlineTextBox",
        "inputTime",
        "keyboard",
        "labelText",
        "layoutTable",
        "layoutTableCell",
        "layoutTableRow",
        "legend",
        "lineBreak",
        "link",
        "list",
        "listBox",
        "listBoxOption",
        "listGrid",
        "listItem",
        "listMarker",
        "log",
        "main",
        "mark",
        "marquee",
        "math",
        "mathMLFraction",
        "mathMLIdentifier",
        "mathMLMath",
        "mathMLMultiscripts",
        "mathMLNoneScript",
        "mathMLNumber",
        "mathMLOperator",
        "mathMLOver",
        "mathMLPrescriptDelimiter",
        "mathMLRoot",
        "mathMLRow",
        "mathMLSquareRoot",
        "mathMLStringLiteral",
        "mathMLSub",
        "mathMLSubSup",
        "mathMLSup",
        "mathMLTable",
        "mathMLTableCell",
        "mathMLTableRow",
        "mathMLText",
        "mathMLUnder",
        "mathMLUnderOver",
        "menu",
        "menuBar",
        "menuItem",
        "menuItemCheckBox",
        "menuItemRadio",
        "menuListOption",
        "menuListPopup",
        "meter",
        "navigation",
        "note",
        "pane",
        "paragraph",
        "pdfActionableHighlight",
        "pdfRoot",
        "pluginObject",
        "popUpButton",
        "portal",
        "preDeprecated",
        "progressIndicator",
        "radioButton",
        "radioGroup",
        "region",
        "rootWebArea",
        "row",
        "rowGroup",
        "rowHeader",
        "ruby",
        "rubyAnnotation",
        "scrollBar",
        "scrollView",
        "search",
        "searchBox",
        "section",
        "slider",
        "spinButton",
        "splitter",
        "staticText",
        "status",
        "strong",
        "subscript",
        "suggestion",
        "superscript",
        "svgRoot",
        "switch",
        "tab",
        "tabList",
        "tabPanel",
        "table",
        "tableHeaderContainer",
        "term",
        "textField",
        "textFieldWithComboBox",
        "time",
        "timer",
        "titleBar",
        "toggleButton",
        "toolbar",
        "tooltip",
        "tree",
        "treeGrid",
        "treeItem",
        "unknown",
        "video",
        "webView",
        "window",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_FindParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 12, false);
        A.store.Enum(ptr + 0, -1);
        A.store.Ref(ptr + 4, undefined);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 12, true);
        A.store.Enum(
          ptr + 0,
          [
            "abbr",
            "alert",
            "alertDialog",
            "application",
            "article",
            "audio",
            "banner",
            "blockquote",
            "button",
            "canvas",
            "caption",
            "caret",
            "cell",
            "checkBox",
            "client",
            "code",
            "colorWell",
            "column",
            "columnHeader",
            "comboBoxGrouping",
            "comboBoxMenuButton",
            "comboBoxSelect",
            "comment",
            "complementary",
            "contentDeletion",
            "contentInsertion",
            "contentInfo",
            "date",
            "dateTime",
            "definition",
            "descriptionList",
            "descriptionListDetail",
            "descriptionListTerm",
            "desktop",
            "details",
            "dialog",
            "directory",
            "disclosureTriangle",
            "docAbstract",
            "docAcknowledgments",
            "docAfterword",
            "docAppendix",
            "docBackLink",
            "docBiblioEntry",
            "docBibliography",
            "docBiblioRef",
            "docChapter",
            "docColophon",
            "docConclusion",
            "docCover",
            "docCredit",
            "docCredits",
            "docDedication",
            "docEndnote",
            "docEndnotes",
            "docEpigraph",
            "docEpilogue",
            "docErrata",
            "docExample",
            "docFootnote",
            "docForeword",
            "docGlossary",
            "docGlossRef",
            "docIndex",
            "docIntroduction",
            "docNoteRef",
            "docNotice",
            "docPageBreak",
            "docPageFooter",
            "docPageHeader",
            "docPageList",
            "docPart",
            "docPreface",
            "docPrologue",
            "docPullquote",
            "docQna",
            "docSubtitle",
            "docTip",
            "docToc",
            "document",
            "embeddedObject",
            "emphasis",
            "feed",
            "figcaption",
            "figure",
            "footer",
            "footerAsNonLandmark",
            "form",
            "genericContainer",
            "graphicsDocument",
            "graphicsObject",
            "graphicsSymbol",
            "grid",
            "group",
            "header",
            "headerAsNonLandmark",
            "heading",
            "iframe",
            "iframePresentational",
            "image",
            "imeCandidate",
            "inlineTextBox",
            "inputTime",
            "keyboard",
            "labelText",
            "layoutTable",
            "layoutTableCell",
            "layoutTableRow",
            "legend",
            "lineBreak",
            "link",
            "list",
            "listBox",
            "listBoxOption",
            "listGrid",
            "listItem",
            "listMarker",
            "log",
            "main",
            "mark",
            "marquee",
            "math",
            "mathMLFraction",
            "mathMLIdentifier",
            "mathMLMath",
            "mathMLMultiscripts",
            "mathMLNoneScript",
            "mathMLNumber",
            "mathMLOperator",
            "mathMLOver",
            "mathMLPrescriptDelimiter",
            "mathMLRoot",
            "mathMLRow",
            "mathMLSquareRoot",
            "mathMLStringLiteral",
            "mathMLSub",
            "mathMLSubSup",
            "mathMLSup",
            "mathMLTable",
            "mathMLTableCell",
            "mathMLTableRow",
            "mathMLText",
            "mathMLUnder",
            "mathMLUnderOver",
            "menu",
            "menuBar",
            "menuItem",
            "menuItemCheckBox",
            "menuItemRadio",
            "menuListOption",
            "menuListPopup",
            "meter",
            "navigation",
            "note",
            "pane",
            "paragraph",
            "pdfActionableHighlight",
            "pdfRoot",
            "pluginObject",
            "popUpButton",
            "portal",
            "preDeprecated",
            "progressIndicator",
            "radioButton",
            "radioGroup",
            "region",
            "rootWebArea",
            "row",
            "rowGroup",
            "rowHeader",
            "ruby",
            "rubyAnnotation",
            "scrollBar",
            "scrollView",
            "search",
            "searchBox",
            "section",
            "slider",
            "spinButton",
            "splitter",
            "staticText",
            "status",
            "strong",
            "subscript",
            "suggestion",
            "superscript",
            "svgRoot",
            "switch",
            "tab",
            "tabList",
            "tabPanel",
            "table",
            "tableHeaderContainer",
            "term",
            "textField",
            "textFieldWithComboBox",
            "time",
            "timer",
            "titleBar",
            "toggleButton",
            "toolbar",
            "tooltip",
            "tree",
            "treeGrid",
            "treeItem",
            "unknown",
            "video",
            "webView",
            "window",
          ].indexOf(x["role"] as string)
        );
        A.store.Ref(ptr + 4, x["state"]);
        A.store.Ref(ptr + 8, x["attributes"]);
      }
    },
    "load_FindParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["role"] = A.load.Enum(ptr + 0, [
        "abbr",
        "alert",
        "alertDialog",
        "application",
        "article",
        "audio",
        "banner",
        "blockquote",
        "button",
        "canvas",
        "caption",
        "caret",
        "cell",
        "checkBox",
        "client",
        "code",
        "colorWell",
        "column",
        "columnHeader",
        "comboBoxGrouping",
        "comboBoxMenuButton",
        "comboBoxSelect",
        "comment",
        "complementary",
        "contentDeletion",
        "contentInsertion",
        "contentInfo",
        "date",
        "dateTime",
        "definition",
        "descriptionList",
        "descriptionListDetail",
        "descriptionListTerm",
        "desktop",
        "details",
        "dialog",
        "directory",
        "disclosureTriangle",
        "docAbstract",
        "docAcknowledgments",
        "docAfterword",
        "docAppendix",
        "docBackLink",
        "docBiblioEntry",
        "docBibliography",
        "docBiblioRef",
        "docChapter",
        "docColophon",
        "docConclusion",
        "docCover",
        "docCredit",
        "docCredits",
        "docDedication",
        "docEndnote",
        "docEndnotes",
        "docEpigraph",
        "docEpilogue",
        "docErrata",
        "docExample",
        "docFootnote",
        "docForeword",
        "docGlossary",
        "docGlossRef",
        "docIndex",
        "docIntroduction",
        "docNoteRef",
        "docNotice",
        "docPageBreak",
        "docPageFooter",
        "docPageHeader",
        "docPageList",
        "docPart",
        "docPreface",
        "docPrologue",
        "docPullquote",
        "docQna",
        "docSubtitle",
        "docTip",
        "docToc",
        "document",
        "embeddedObject",
        "emphasis",
        "feed",
        "figcaption",
        "figure",
        "footer",
        "footerAsNonLandmark",
        "form",
        "genericContainer",
        "graphicsDocument",
        "graphicsObject",
        "graphicsSymbol",
        "grid",
        "group",
        "header",
        "headerAsNonLandmark",
        "heading",
        "iframe",
        "iframePresentational",
        "image",
        "imeCandidate",
        "inlineTextBox",
        "inputTime",
        "keyboard",
        "labelText",
        "layoutTable",
        "layoutTableCell",
        "layoutTableRow",
        "legend",
        "lineBreak",
        "link",
        "list",
        "listBox",
        "listBoxOption",
        "listGrid",
        "listItem",
        "listMarker",
        "log",
        "main",
        "mark",
        "marquee",
        "math",
        "mathMLFraction",
        "mathMLIdentifier",
        "mathMLMath",
        "mathMLMultiscripts",
        "mathMLNoneScript",
        "mathMLNumber",
        "mathMLOperator",
        "mathMLOver",
        "mathMLPrescriptDelimiter",
        "mathMLRoot",
        "mathMLRow",
        "mathMLSquareRoot",
        "mathMLStringLiteral",
        "mathMLSub",
        "mathMLSubSup",
        "mathMLSup",
        "mathMLTable",
        "mathMLTableCell",
        "mathMLTableRow",
        "mathMLText",
        "mathMLUnder",
        "mathMLUnderOver",
        "menu",
        "menuBar",
        "menuItem",
        "menuItemCheckBox",
        "menuItemRadio",
        "menuListOption",
        "menuListPopup",
        "meter",
        "navigation",
        "note",
        "pane",
        "paragraph",
        "pdfActionableHighlight",
        "pdfRoot",
        "pluginObject",
        "popUpButton",
        "portal",
        "preDeprecated",
        "progressIndicator",
        "radioButton",
        "radioGroup",
        "region",
        "rootWebArea",
        "row",
        "rowGroup",
        "rowHeader",
        "ruby",
        "rubyAnnotation",
        "scrollBar",
        "scrollView",
        "search",
        "searchBox",
        "section",
        "slider",
        "spinButton",
        "splitter",
        "staticText",
        "status",
        "strong",
        "subscript",
        "suggestion",
        "superscript",
        "svgRoot",
        "switch",
        "tab",
        "tabList",
        "tabPanel",
        "table",
        "tableHeaderContainer",
        "term",
        "textField",
        "textFieldWithComboBox",
        "time",
        "timer",
        "titleBar",
        "toggleButton",
        "toolbar",
        "tooltip",
        "tree",
        "treeGrid",
        "treeItem",
        "unknown",
        "video",
        "webView",
        "window",
      ]);
      x["state"] = A.load.Ref(ptr + 4, undefined);
      x["attributes"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },

    "has_AutomationPosition_IsNullPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isNullPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsNullPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isNullPosition);
    },

    "call_AutomationPosition_IsNullPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isNullPosition();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsNullPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isNullPosition();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_IsTreePosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isTreePosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsTreePosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isTreePosition);
    },

    "call_AutomationPosition_IsTreePosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isTreePosition();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsTreePosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isTreePosition();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_IsTextPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isTextPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsTextPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isTextPosition);
    },

    "call_AutomationPosition_IsTextPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isTextPosition();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsTextPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isTextPosition();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_IsLeafTextPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isLeafTextPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsLeafTextPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isLeafTextPosition);
    },

    "call_AutomationPosition_IsLeafTextPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isLeafTextPosition();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsLeafTextPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isLeafTextPosition();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtStartOfAnchor": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atStartOfAnchor" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtStartOfAnchor": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atStartOfAnchor);
    },

    "call_AutomationPosition_AtStartOfAnchor": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atStartOfAnchor();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtStartOfAnchor": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atStartOfAnchor();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtEndOfAnchor": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atEndOfAnchor" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtEndOfAnchor": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atEndOfAnchor);
    },

    "call_AutomationPosition_AtEndOfAnchor": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atEndOfAnchor();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtEndOfAnchor": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atEndOfAnchor();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtStartOfWord": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atStartOfWord" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtStartOfWord": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atStartOfWord);
    },

    "call_AutomationPosition_AtStartOfWord": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atStartOfWord();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtStartOfWord": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atStartOfWord();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtEndOfWord": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atEndOfWord" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtEndOfWord": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atEndOfWord);
    },

    "call_AutomationPosition_AtEndOfWord": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atEndOfWord();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtEndOfWord": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atEndOfWord();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtStartOfLine": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atStartOfLine" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtStartOfLine": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atStartOfLine);
    },

    "call_AutomationPosition_AtStartOfLine": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atStartOfLine();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtStartOfLine": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atStartOfLine();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtEndOfLine": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atEndOfLine" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtEndOfLine": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atEndOfLine);
    },

    "call_AutomationPosition_AtEndOfLine": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atEndOfLine();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtEndOfLine": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atEndOfLine();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtStartOfParagraph": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atStartOfParagraph" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtStartOfParagraph": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atStartOfParagraph);
    },

    "call_AutomationPosition_AtStartOfParagraph": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atStartOfParagraph();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtStartOfParagraph": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atStartOfParagraph();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtEndOfParagraph": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atEndOfParagraph" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtEndOfParagraph": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atEndOfParagraph);
    },

    "call_AutomationPosition_AtEndOfParagraph": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atEndOfParagraph();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtEndOfParagraph": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atEndOfParagraph();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtStartOfPage": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atStartOfPage" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtStartOfPage": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atStartOfPage);
    },

    "call_AutomationPosition_AtStartOfPage": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atStartOfPage();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtStartOfPage": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atStartOfPage();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtEndOfPage": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atEndOfPage" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtEndOfPage": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atEndOfPage);
    },

    "call_AutomationPosition_AtEndOfPage": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atEndOfPage();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtEndOfPage": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atEndOfPage();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtStartOfFormat": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atStartOfFormat" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtStartOfFormat": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atStartOfFormat);
    },

    "call_AutomationPosition_AtStartOfFormat": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atStartOfFormat();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtStartOfFormat": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atStartOfFormat();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtEndOfFormat": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atEndOfFormat" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtEndOfFormat": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atEndOfFormat);
    },

    "call_AutomationPosition_AtEndOfFormat": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atEndOfFormat();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtEndOfFormat": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atEndOfFormat();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtStartOfDocument": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atStartOfDocument" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtStartOfDocument": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atStartOfDocument);
    },

    "call_AutomationPosition_AtStartOfDocument": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atStartOfDocument();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtStartOfDocument": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atStartOfDocument();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AtEndOfDocument": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "atEndOfDocument" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AtEndOfDocument": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).atEndOfDocument);
    },

    "call_AutomationPosition_AtEndOfDocument": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.atEndOfDocument();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_AtEndOfDocument": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.atEndOfDocument();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AsTreePosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "asTreePosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AsTreePosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).asTreePosition);
    },

    "call_AutomationPosition_AsTreePosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.asTreePosition();
    },
    "try_AutomationPosition_AsTreePosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.asTreePosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AsTextPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "asTextPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AsTextPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).asTextPosition);
    },

    "call_AutomationPosition_AsTextPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.asTextPosition();
    },
    "try_AutomationPosition_AsTextPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.asTextPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_AsLeafTextPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "asLeafTextPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_AsLeafTextPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).asLeafTextPosition);
    },

    "call_AutomationPosition_AsLeafTextPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.asLeafTextPosition();
    },
    "try_AutomationPosition_AsLeafTextPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.asLeafTextPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPositionAtStartOfAnchor": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPositionAtStartOfAnchor" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPositionAtStartOfAnchor": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPositionAtStartOfAnchor);
    },

    "call_AutomationPosition_MoveToPositionAtStartOfAnchor": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPositionAtStartOfAnchor();
    },
    "try_AutomationPosition_MoveToPositionAtStartOfAnchor": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPositionAtStartOfAnchor();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPositionAtEndOfAnchor": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPositionAtEndOfAnchor" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPositionAtEndOfAnchor": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPositionAtEndOfAnchor);
    },

    "call_AutomationPosition_MoveToPositionAtEndOfAnchor": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPositionAtEndOfAnchor();
    },
    "try_AutomationPosition_MoveToPositionAtEndOfAnchor": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPositionAtEndOfAnchor();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPositionAtStartOfDocument": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPositionAtStartOfDocument" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPositionAtStartOfDocument": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPositionAtStartOfDocument);
    },

    "call_AutomationPosition_MoveToPositionAtStartOfDocument": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPositionAtStartOfDocument();
    },
    "try_AutomationPosition_MoveToPositionAtStartOfDocument": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPositionAtStartOfDocument();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPositionAtEndOfDocument": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPositionAtEndOfDocument" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPositionAtEndOfDocument": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPositionAtEndOfDocument);
    },

    "call_AutomationPosition_MoveToPositionAtEndOfDocument": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPositionAtEndOfDocument();
    },
    "try_AutomationPosition_MoveToPositionAtEndOfDocument": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPositionAtEndOfDocument();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToParentPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToParentPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToParentPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToParentPosition);
    },

    "call_AutomationPosition_MoveToParentPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToParentPosition();
    },
    "try_AutomationPosition_MoveToParentPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToParentPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextLeafTreePosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextLeafTreePosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextLeafTreePosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextLeafTreePosition);
    },

    "call_AutomationPosition_MoveToNextLeafTreePosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextLeafTreePosition();
    },
    "try_AutomationPosition_MoveToNextLeafTreePosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextLeafTreePosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousLeafTreePosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousLeafTreePosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousLeafTreePosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousLeafTreePosition);
    },

    "call_AutomationPosition_MoveToPreviousLeafTreePosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousLeafTreePosition();
    },
    "try_AutomationPosition_MoveToPreviousLeafTreePosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousLeafTreePosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextLeafTextPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextLeafTextPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextLeafTextPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextLeafTextPosition);
    },

    "call_AutomationPosition_MoveToNextLeafTextPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextLeafTextPosition();
    },
    "try_AutomationPosition_MoveToNextLeafTextPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextLeafTextPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousLeafTextPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousLeafTextPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousLeafTextPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousLeafTextPosition);
    },

    "call_AutomationPosition_MoveToPreviousLeafTextPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousLeafTextPosition();
    },
    "try_AutomationPosition_MoveToPreviousLeafTextPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousLeafTextPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextCharacterPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextCharacterPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextCharacterPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextCharacterPosition);
    },

    "call_AutomationPosition_MoveToNextCharacterPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextCharacterPosition();
    },
    "try_AutomationPosition_MoveToNextCharacterPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextCharacterPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousCharacterPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousCharacterPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousCharacterPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousCharacterPosition);
    },

    "call_AutomationPosition_MoveToPreviousCharacterPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousCharacterPosition();
    },
    "try_AutomationPosition_MoveToPreviousCharacterPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousCharacterPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextWordStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextWordStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextWordStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextWordStartPosition);
    },

    "call_AutomationPosition_MoveToNextWordStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextWordStartPosition();
    },
    "try_AutomationPosition_MoveToNextWordStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextWordStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousWordStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousWordStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousWordStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousWordStartPosition);
    },

    "call_AutomationPosition_MoveToPreviousWordStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousWordStartPosition();
    },
    "try_AutomationPosition_MoveToPreviousWordStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousWordStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextWordEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextWordEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextWordEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextWordEndPosition);
    },

    "call_AutomationPosition_MoveToNextWordEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextWordEndPosition();
    },
    "try_AutomationPosition_MoveToNextWordEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextWordEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousWordEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousWordEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousWordEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousWordEndPosition);
    },

    "call_AutomationPosition_MoveToPreviousWordEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousWordEndPosition();
    },
    "try_AutomationPosition_MoveToPreviousWordEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousWordEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextLineStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextLineStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextLineStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextLineStartPosition);
    },

    "call_AutomationPosition_MoveToNextLineStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextLineStartPosition();
    },
    "try_AutomationPosition_MoveToNextLineStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextLineStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousLineStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousLineStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousLineStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousLineStartPosition);
    },

    "call_AutomationPosition_MoveToPreviousLineStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousLineStartPosition();
    },
    "try_AutomationPosition_MoveToPreviousLineStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousLineStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextLineEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextLineEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextLineEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextLineEndPosition);
    },

    "call_AutomationPosition_MoveToNextLineEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextLineEndPosition();
    },
    "try_AutomationPosition_MoveToNextLineEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextLineEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousLineEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousLineEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousLineEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousLineEndPosition);
    },

    "call_AutomationPosition_MoveToPreviousLineEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousLineEndPosition();
    },
    "try_AutomationPosition_MoveToPreviousLineEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousLineEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextFormatStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextFormatStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextFormatStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextFormatStartPosition);
    },

    "call_AutomationPosition_MoveToNextFormatStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextFormatStartPosition();
    },
    "try_AutomationPosition_MoveToNextFormatStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextFormatStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousFormatStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousFormatStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousFormatStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousFormatStartPosition);
    },

    "call_AutomationPosition_MoveToPreviousFormatStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousFormatStartPosition();
    },
    "try_AutomationPosition_MoveToPreviousFormatStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousFormatStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextFormatEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextFormatEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextFormatEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextFormatEndPosition);
    },

    "call_AutomationPosition_MoveToNextFormatEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextFormatEndPosition();
    },
    "try_AutomationPosition_MoveToNextFormatEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextFormatEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousFormatEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousFormatEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousFormatEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousFormatEndPosition);
    },

    "call_AutomationPosition_MoveToPreviousFormatEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousFormatEndPosition();
    },
    "try_AutomationPosition_MoveToPreviousFormatEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousFormatEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextParagraphStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextParagraphStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextParagraphStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextParagraphStartPosition);
    },

    "call_AutomationPosition_MoveToNextParagraphStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextParagraphStartPosition();
    },
    "try_AutomationPosition_MoveToNextParagraphStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextParagraphStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousParagraphStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousParagraphStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousParagraphStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousParagraphStartPosition);
    },

    "call_AutomationPosition_MoveToPreviousParagraphStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousParagraphStartPosition();
    },
    "try_AutomationPosition_MoveToPreviousParagraphStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousParagraphStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextParagraphEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextParagraphEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextParagraphEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextParagraphEndPosition);
    },

    "call_AutomationPosition_MoveToNextParagraphEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextParagraphEndPosition();
    },
    "try_AutomationPosition_MoveToNextParagraphEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextParagraphEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousParagraphEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousParagraphEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousParagraphEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousParagraphEndPosition);
    },

    "call_AutomationPosition_MoveToPreviousParagraphEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousParagraphEndPosition();
    },
    "try_AutomationPosition_MoveToPreviousParagraphEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousParagraphEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextPageStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextPageStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextPageStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextPageStartPosition);
    },

    "call_AutomationPosition_MoveToNextPageStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextPageStartPosition();
    },
    "try_AutomationPosition_MoveToNextPageStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextPageStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousPageStartPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousPageStartPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousPageStartPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousPageStartPosition);
    },

    "call_AutomationPosition_MoveToPreviousPageStartPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousPageStartPosition();
    },
    "try_AutomationPosition_MoveToPreviousPageStartPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousPageStartPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextPageEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextPageEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextPageEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextPageEndPosition);
    },

    "call_AutomationPosition_MoveToNextPageEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextPageEndPosition();
    },
    "try_AutomationPosition_MoveToNextPageEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextPageEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousPageEndPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousPageEndPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousPageEndPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousPageEndPosition);
    },

    "call_AutomationPosition_MoveToPreviousPageEndPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousPageEndPosition();
    },
    "try_AutomationPosition_MoveToPreviousPageEndPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousPageEndPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToNextAnchorPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToNextAnchorPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToNextAnchorPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToNextAnchorPosition);
    },

    "call_AutomationPosition_MoveToNextAnchorPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToNextAnchorPosition();
    },
    "try_AutomationPosition_MoveToNextAnchorPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToNextAnchorPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MoveToPreviousAnchorPosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "moveToPreviousAnchorPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MoveToPreviousAnchorPosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).moveToPreviousAnchorPosition);
    },

    "call_AutomationPosition_MoveToPreviousAnchorPosition": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.moveToPreviousAnchorPosition();
    },
    "try_AutomationPosition_MoveToPreviousAnchorPosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.moveToPreviousAnchorPosition();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_MaxTextOffset": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "maxTextOffset" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_MaxTextOffset": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).maxTextOffset);
    },

    "call_AutomationPosition_MaxTextOffset": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.maxTextOffset();
      A.store.Int32(retPtr, _ret);
    },
    "try_AutomationPosition_MaxTextOffset": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.maxTextOffset();
        A.store.Int32(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_IsInLineBreak": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isInLineBreak" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsInLineBreak": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isInLineBreak);
    },

    "call_AutomationPosition_IsInLineBreak": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isInLineBreak();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsInLineBreak": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isInLineBreak();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_IsInTextObject": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isInTextObject" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsInTextObject": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isInTextObject);
    },

    "call_AutomationPosition_IsInTextObject": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isInTextObject();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsInTextObject": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isInTextObject();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_IsInWhiteSpace": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isInWhiteSpace" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsInWhiteSpace": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isInWhiteSpace);
    },

    "call_AutomationPosition_IsInWhiteSpace": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isInWhiteSpace();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsInWhiteSpace": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isInWhiteSpace();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_IsValid": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "isValid" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_IsValid": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).isValid);
    },

    "call_AutomationPosition_IsValid": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.isValid();
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationPosition_IsValid": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.isValid();
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationPosition_GetText": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getText" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationPosition_GetText": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getText);
    },

    "call_AutomationPosition_GetText": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getText();
      A.store.Ref(retPtr, _ret);
    },
    "try_AutomationPosition_GetText": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getText();
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_AutomationPosition_Node": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "node" in thiz) {
        const val = thiz.node;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationPosition_Node": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "node", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationPosition_ChildIndex": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "childIndex" in thiz) {
        const val = thiz.childIndex;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationPosition_ChildIndex": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "childIndex", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationPosition_TextOffset": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "textOffset" in thiz) {
        const val = thiz.textOffset;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationPosition_TextOffset": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "textOffset", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationPosition_Affinity": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "affinity" in thiz) {
        const val = thiz.affinity;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationPosition_Affinity": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "affinity", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "constof_PositionType": (ref: heap.Ref<string>): number => {
      const idx = ["null", "text", "tree"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_NameFromType": (ref: heap.Ref<string>): number => {
      const idx = [
        "attribute",
        "attributeExplicitlyEmpty",
        "caption",
        "contents",
        "placeholder",
        "popoverAttribute",
        "relatedElement",
        "title",
        "value",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_CustomAction": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 9, false);
        A.store.Bool(ptr + 8, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Ref(ptr + 4, undefined);
      } else {
        A.store.Bool(ptr + 9, true);
        A.store.Bool(ptr + 8, "id" in x ? true : false);
        A.store.Int32(ptr + 0, x["id"] === undefined ? 0 : (x["id"] as number));
        A.store.Ref(ptr + 4, x["description"]);
      }
    },
    "load_CustomAction": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 8)) {
        x["id"] = A.load.Int32(ptr + 0);
      } else {
        delete x["id"];
      }
      x["description"] = A.load.Ref(ptr + 4, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_DefaultActionVerb": (ref: heap.Ref<string>): number => {
      const idx = ["activate", "check", "click", "clickAncestor", "jump", "open", "press", "select", "uncheck"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },

    "store_Marker": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 14, false);
        A.store.Bool(ptr + 12, false);
        A.store.Int32(ptr + 0, 0);
        A.store.Bool(ptr + 13, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
      } else {
        A.store.Bool(ptr + 14, true);
        A.store.Bool(ptr + 12, "startOffset" in x ? true : false);
        A.store.Int32(ptr + 0, x["startOffset"] === undefined ? 0 : (x["startOffset"] as number));
        A.store.Bool(ptr + 13, "endOffset" in x ? true : false);
        A.store.Int32(ptr + 4, x["endOffset"] === undefined ? 0 : (x["endOffset"] as number));
        A.store.Ref(ptr + 8, x["flags"]);
      }
    },
    "load_Marker": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      if (A.load.Bool(ptr + 12)) {
        x["startOffset"] = A.load.Int32(ptr + 0);
      } else {
        delete x["startOffset"];
      }
      if (A.load.Bool(ptr + 13)) {
        x["endOffset"] = A.load.Int32(ptr + 4);
      } else {
        delete x["endOffset"];
      }
      x["flags"] = A.load.Ref(ptr + 8, undefined);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_HasPopup": (ref: heap.Ref<string>): number => {
      const idx = ["false", "true", "menu", "listbox", "tree", "grid", "dialog"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_AriaCurrentState": (ref: heap.Ref<string>): number => {
      const idx = ["false", "true", "page", "step", "location", "date", "time"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_InvalidState": (ref: heap.Ref<string>): number => {
      const idx = ["false", "true"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_SortDirectionType": (ref: heap.Ref<string>): number => {
      const idx = ["unsorted", "ascending", "descending", "other"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "has_AutomationNode_BoundsForRange": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "boundsForRange" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_BoundsForRange": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).boundsForRange);
    },

    "call_AutomationNode_BoundsForRange": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      startIndex: number,
      endIndex: number,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.boundsForRange(startIndex, endIndex, A.H.get<object>(callback));
    },
    "try_AutomationNode_BoundsForRange": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      startIndex: number,
      endIndex: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.boundsForRange(startIndex, endIndex, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_UnclippedBoundsForRange": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "unclippedBoundsForRange" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_UnclippedBoundsForRange": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).unclippedBoundsForRange);
    },

    "call_AutomationNode_UnclippedBoundsForRange": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      startIndex: number,
      endIndex: number,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.unclippedBoundsForRange(startIndex, endIndex, A.H.get<object>(callback));
    },
    "try_AutomationNode_UnclippedBoundsForRange": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      startIndex: number,
      endIndex: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.unclippedBoundsForRange(startIndex, endIndex, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_DoDefault": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "doDefault" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_DoDefault": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).doDefault);
    },

    "call_AutomationNode_DoDefault": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.doDefault();
    },
    "try_AutomationNode_DoDefault": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.doDefault();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_Focus": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "focus" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_Focus": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).focus);
    },

    "call_AutomationNode_Focus": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.focus();
    },
    "try_AutomationNode_Focus": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.focus();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_GetImageData": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getImageData" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_GetImageData": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getImageData);
    },

    "call_AutomationNode_GetImageData": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      maxWidth: number,
      maxHeight: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getImageData(maxWidth, maxHeight);
    },
    "try_AutomationNode_GetImageData": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      maxWidth: number,
      maxHeight: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getImageData(maxWidth, maxHeight);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_HitTest": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "hitTest" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_HitTest": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).hitTest);
    },

    "call_AutomationNode_HitTest": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      x: number,
      y: number,
      eventToFire: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.hitTest(
        x,
        y,
        eventToFire > 0 && eventToFire <= 119
          ? [
              "accessKeyChanged",
              "activeDescendantChanged",
              "alert",
              "ariaAttributeChangedDeprecated",
              "ariaCurrentChanged",
              "atomicChanged",
              "autoCompleteChanged",
              "autocorrectionOccured",
              "autofillAvailabilityChanged",
              "blur",
              "busyChanged",
              "caretBoundsChanged",
              "checkedStateChanged",
              "checkedStateDescriptionChanged",
              "childrenChanged",
              "classNameChanged",
              "clicked",
              "collapsed",
              "controlsChanged",
              "detailsChanged",
              "describedByChanged",
              "descriptionChanged",
              "documentSelectionChanged",
              "documentTitleChanged",
              "dropeffectChanged",
              "editableTextChanged",
              "enabledChanged",
              "endOfTest",
              "expanded",
              "expandedChanged",
              "flowFromChanged",
              "flowToChanged",
              "focus",
              "focusAfterMenuClose",
              "focusChanged",
              "focusContext",
              "grabbedChanged",
              "haspopupChanged",
              "hide",
              "hierarchicalLevelChanged",
              "hitTestResult",
              "hover",
              "ignoredChanged",
              "imageAnnotationChanged",
              "imageFrameUpdated",
              "invalidStatusChanged",
              "keyShortcutsChanged",
              "labeledByChanged",
              "languageChanged",
              "layoutComplete",
              "layoutInvalidated",
              "liveRegionChanged",
              "liveRegionCreated",
              "liveRegionNodeChanged",
              "liveRelevantChanged",
              "liveStatusChanged",
              "loadComplete",
              "loadStart",
              "locationChanged",
              "mediaStartedPlaying",
              "mediaStoppedPlaying",
              "menuEnd",
              "menuItemSelected",
              "menuListValueChanged",
              "menuPopupEnd",
              "menuPopupStart",
              "menuStart",
              "mouseCanceled",
              "mouseDragged",
              "mouseMoved",
              "mousePressed",
              "mouseReleased",
              "multilineStateChanged",
              "multiselectableStateChanged",
              "nameChanged",
              "objectAttributeChanged",
              "orientationChanged",
              "otherAttributeChanged",
              "parentChanged",
              "placeholderChanged",
              "portalActivated",
              "positionInSetChanged",
              "rangeValueChanged",
              "rangeValueMaxChanged",
              "rangeValueMinChanged",
              "rangeValueStepChanged",
              "readonlyChanged",
              "relatedNodeChanged",
              "requiredStateChanged",
              "roleChanged",
              "rowCollapsed",
              "rowCountChanged",
              "rowExpanded",
              "scrollHorizontalPositionChanged",
              "scrollPositionChanged",
              "scrollVerticalPositionChanged",
              "scrolledToAnchor",
              "selectedChanged",
              "selectedChildrenChanged",
              "selectedValueChanged",
              "selection",
              "selectionAdd",
              "selectionRemove",
              "setSizeChanged",
              "show",
              "sortChanged",
              "stateChanged",
              "subtreeCreated",
              "textAttributeChanged",
              "textSelectionChanged",
              "textChanged",
              "tooltipClosed",
              "tooltipOpened",
              "treeChanged",
              "valueInTextFieldChanged",
              "valueChanged",
              "windowActivated",
              "windowDeactivated",
              "windowVisibilityChanged",
            ][eventToFire - 1]
          : undefined
      );
    },
    "try_AutomationNode_HitTest": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      x: number,
      y: number,
      eventToFire: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.hitTest(
          x,
          y,
          eventToFire > 0 && eventToFire <= 119
            ? [
                "accessKeyChanged",
                "activeDescendantChanged",
                "alert",
                "ariaAttributeChangedDeprecated",
                "ariaCurrentChanged",
                "atomicChanged",
                "autoCompleteChanged",
                "autocorrectionOccured",
                "autofillAvailabilityChanged",
                "blur",
                "busyChanged",
                "caretBoundsChanged",
                "checkedStateChanged",
                "checkedStateDescriptionChanged",
                "childrenChanged",
                "classNameChanged",
                "clicked",
                "collapsed",
                "controlsChanged",
                "detailsChanged",
                "describedByChanged",
                "descriptionChanged",
                "documentSelectionChanged",
                "documentTitleChanged",
                "dropeffectChanged",
                "editableTextChanged",
                "enabledChanged",
                "endOfTest",
                "expanded",
                "expandedChanged",
                "flowFromChanged",
                "flowToChanged",
                "focus",
                "focusAfterMenuClose",
                "focusChanged",
                "focusContext",
                "grabbedChanged",
                "haspopupChanged",
                "hide",
                "hierarchicalLevelChanged",
                "hitTestResult",
                "hover",
                "ignoredChanged",
                "imageAnnotationChanged",
                "imageFrameUpdated",
                "invalidStatusChanged",
                "keyShortcutsChanged",
                "labeledByChanged",
                "languageChanged",
                "layoutComplete",
                "layoutInvalidated",
                "liveRegionChanged",
                "liveRegionCreated",
                "liveRegionNodeChanged",
                "liveRelevantChanged",
                "liveStatusChanged",
                "loadComplete",
                "loadStart",
                "locationChanged",
                "mediaStartedPlaying",
                "mediaStoppedPlaying",
                "menuEnd",
                "menuItemSelected",
                "menuListValueChanged",
                "menuPopupEnd",
                "menuPopupStart",
                "menuStart",
                "mouseCanceled",
                "mouseDragged",
                "mouseMoved",
                "mousePressed",
                "mouseReleased",
                "multilineStateChanged",
                "multiselectableStateChanged",
                "nameChanged",
                "objectAttributeChanged",
                "orientationChanged",
                "otherAttributeChanged",
                "parentChanged",
                "placeholderChanged",
                "portalActivated",
                "positionInSetChanged",
                "rangeValueChanged",
                "rangeValueMaxChanged",
                "rangeValueMinChanged",
                "rangeValueStepChanged",
                "readonlyChanged",
                "relatedNodeChanged",
                "requiredStateChanged",
                "roleChanged",
                "rowCollapsed",
                "rowCountChanged",
                "rowExpanded",
                "scrollHorizontalPositionChanged",
                "scrollPositionChanged",
                "scrollVerticalPositionChanged",
                "scrolledToAnchor",
                "selectedChanged",
                "selectedChildrenChanged",
                "selectedValueChanged",
                "selection",
                "selectionAdd",
                "selectionRemove",
                "setSizeChanged",
                "show",
                "sortChanged",
                "stateChanged",
                "subtreeCreated",
                "textAttributeChanged",
                "textSelectionChanged",
                "textChanged",
                "tooltipClosed",
                "tooltipOpened",
                "treeChanged",
                "valueInTextFieldChanged",
                "valueChanged",
                "windowActivated",
                "windowDeactivated",
                "windowVisibilityChanged",
              ][eventToFire - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_HitTestWithReply": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "hitTestWithReply" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_HitTestWithReply": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).hitTestWithReply);
    },

    "call_AutomationNode_HitTestWithReply": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      x: number,
      y: number,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.hitTestWithReply(x, y, A.H.get<object>(callback));
    },
    "try_AutomationNode_HitTestWithReply": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      x: number,
      y: number,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.hitTestWithReply(x, y, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_MakeVisible": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "makeVisible" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_MakeVisible": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).makeVisible);
    },

    "call_AutomationNode_MakeVisible": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.makeVisible();
    },
    "try_AutomationNode_MakeVisible": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.makeVisible();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_PerformCustomAction": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "performCustomAction" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_PerformCustomAction": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).performCustomAction);
    },

    "call_AutomationNode_PerformCustomAction": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      customActionId: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.performCustomAction(customActionId);
    },
    "try_AutomationNode_PerformCustomAction": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      customActionId: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.performCustomAction(customActionId);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_PerformStandardAction": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "performStandardAction" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_PerformStandardAction": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).performStandardAction);
    },

    "call_AutomationNode_PerformStandardAction": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      actionType: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.performStandardAction(
        actionType > 0 && actionType <= 39
          ? [
              "annotatePageImages",
              "blur",
              "clearAccessibilityFocus",
              "collapse",
              "customAction",
              "decrement",
              "doDefault",
              "expand",
              "focus",
              "getImageData",
              "getTextLocation",
              "hideTooltip",
              "hitTest",
              "increment",
              "internalInvalidateTree",
              "loadInlineTextBoxes",
              "longClick",
              "replaceSelectedText",
              "resumeMedia",
              "scrollBackward",
              "scrollDown",
              "scrollForward",
              "scrollLeft",
              "scrollRight",
              "scrollUp",
              "scrollToMakeVisible",
              "scrollToPoint",
              "scrollToPositionAtRowColumn",
              "setAccessibilityFocus",
              "setScrollOffset",
              "setSelection",
              "setSequentialFocusNavigationStartingPoint",
              "setValue",
              "showContextMenu",
              "signalEndOfTest",
              "showTooltip",
              "startDuckingMedia",
              "stopDuckingMedia",
              "suspendMedia",
            ][actionType - 1]
          : undefined
      );
    },
    "try_AutomationNode_PerformStandardAction": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      actionType: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.performStandardAction(
          actionType > 0 && actionType <= 39
            ? [
                "annotatePageImages",
                "blur",
                "clearAccessibilityFocus",
                "collapse",
                "customAction",
                "decrement",
                "doDefault",
                "expand",
                "focus",
                "getImageData",
                "getTextLocation",
                "hideTooltip",
                "hitTest",
                "increment",
                "internalInvalidateTree",
                "loadInlineTextBoxes",
                "longClick",
                "replaceSelectedText",
                "resumeMedia",
                "scrollBackward",
                "scrollDown",
                "scrollForward",
                "scrollLeft",
                "scrollRight",
                "scrollUp",
                "scrollToMakeVisible",
                "scrollToPoint",
                "scrollToPositionAtRowColumn",
                "setAccessibilityFocus",
                "setScrollOffset",
                "setSelection",
                "setSequentialFocusNavigationStartingPoint",
                "setValue",
                "showContextMenu",
                "signalEndOfTest",
                "showTooltip",
                "startDuckingMedia",
                "stopDuckingMedia",
                "suspendMedia",
              ][actionType - 1]
            : undefined
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ReplaceSelectedText": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "replaceSelectedText" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ReplaceSelectedText": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).replaceSelectedText);
    },

    "call_AutomationNode_ReplaceSelectedText": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      value: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.replaceSelectedText(A.H.get<object>(value));
    },
    "try_AutomationNode_ReplaceSelectedText": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.replaceSelectedText(A.H.get<object>(value));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_SetAccessibilityFocus": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setAccessibilityFocus" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_SetAccessibilityFocus": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setAccessibilityFocus);
    },

    "call_AutomationNode_SetAccessibilityFocus": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setAccessibilityFocus();
    },
    "try_AutomationNode_SetAccessibilityFocus": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setAccessibilityFocus();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_SetSelection": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setSelection" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_SetSelection": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setSelection);
    },

    "call_AutomationNode_SetSelection": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      startIndex: number,
      endIndex: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setSelection(startIndex, endIndex);
    },
    "try_AutomationNode_SetSelection": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      startIndex: number,
      endIndex: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setSelection(startIndex, endIndex);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_SetSequentialFocusNavigationStartingPoint": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setSequentialFocusNavigationStartingPoint" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_SetSequentialFocusNavigationStartingPoint": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setSequentialFocusNavigationStartingPoint);
    },

    "call_AutomationNode_SetSequentialFocusNavigationStartingPoint": (
      self: heap.Ref<object>,
      retPtr: Pointer
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setSequentialFocusNavigationStartingPoint();
    },
    "try_AutomationNode_SetSequentialFocusNavigationStartingPoint": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setSequentialFocusNavigationStartingPoint();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_SetValue": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setValue" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_SetValue": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setValue);
    },

    "call_AutomationNode_SetValue": (self: heap.Ref<object>, retPtr: Pointer, value: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setValue(A.H.get<object>(value));
    },
    "try_AutomationNode_SetValue": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      value: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setValue(A.H.get<object>(value));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ShowContextMenu": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "showContextMenu" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ShowContextMenu": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).showContextMenu);
    },

    "call_AutomationNode_ShowContextMenu": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.showContextMenu();
    },
    "try_AutomationNode_ShowContextMenu": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.showContextMenu();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ResumeMedia": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "resumeMedia" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ResumeMedia": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).resumeMedia);
    },

    "call_AutomationNode_ResumeMedia": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.resumeMedia();
    },
    "try_AutomationNode_ResumeMedia": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.resumeMedia();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_StartDuckingMedia": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "startDuckingMedia" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_StartDuckingMedia": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).startDuckingMedia);
    },

    "call_AutomationNode_StartDuckingMedia": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.startDuckingMedia();
    },
    "try_AutomationNode_StartDuckingMedia": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.startDuckingMedia();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_StopDuckingMedia": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "stopDuckingMedia" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_StopDuckingMedia": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).stopDuckingMedia);
    },

    "call_AutomationNode_StopDuckingMedia": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.stopDuckingMedia();
    },
    "try_AutomationNode_StopDuckingMedia": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.stopDuckingMedia();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_SuspendMedia": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "suspendMedia" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_SuspendMedia": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).suspendMedia);
    },

    "call_AutomationNode_SuspendMedia": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.suspendMedia();
    },
    "try_AutomationNode_SuspendMedia": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.suspendMedia();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_LongClick": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "longClick" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_LongClick": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).longClick);
    },

    "call_AutomationNode_LongClick": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.longClick();
    },
    "try_AutomationNode_LongClick": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.longClick();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollBackward": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollBackward" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollBackward": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollBackward);
    },

    "call_AutomationNode_ScrollBackward": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollBackward(A.H.get<object>(callback));
    },
    "try_AutomationNode_ScrollBackward": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollBackward(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollBackward1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollBackward" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollBackward1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollBackward);
    },

    "call_AutomationNode_ScrollBackward1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollBackward();
    },
    "try_AutomationNode_ScrollBackward1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollBackward();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollForward": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollForward" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollForward": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollForward);
    },

    "call_AutomationNode_ScrollForward": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      callback: heap.Ref<object>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollForward(A.H.get<object>(callback));
    },
    "try_AutomationNode_ScrollForward": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollForward(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollForward1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollForward" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollForward1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollForward);
    },

    "call_AutomationNode_ScrollForward1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollForward();
    },
    "try_AutomationNode_ScrollForward1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollForward();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollUp": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollUp" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollUp": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollUp);
    },

    "call_AutomationNode_ScrollUp": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollUp(A.H.get<object>(callback));
    },
    "try_AutomationNode_ScrollUp": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollUp(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollUp1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollUp" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollUp1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollUp);
    },

    "call_AutomationNode_ScrollUp1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollUp();
    },
    "try_AutomationNode_ScrollUp1": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollUp();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollDown": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollDown" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollDown": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollDown);
    },

    "call_AutomationNode_ScrollDown": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollDown(A.H.get<object>(callback));
    },
    "try_AutomationNode_ScrollDown": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollDown(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollDown1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollDown" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollDown1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollDown);
    },

    "call_AutomationNode_ScrollDown1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollDown();
    },
    "try_AutomationNode_ScrollDown1": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollDown();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollLeft": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollLeft" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollLeft": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollLeft);
    },

    "call_AutomationNode_ScrollLeft": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollLeft(A.H.get<object>(callback));
    },
    "try_AutomationNode_ScrollLeft": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollLeft(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollLeft1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollLeft" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollLeft1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollLeft);
    },

    "call_AutomationNode_ScrollLeft1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollLeft();
    },
    "try_AutomationNode_ScrollLeft1": (self: heap.Ref<object>, retPtr: Pointer, errPtr: Pointer): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollLeft();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollRight": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollRight" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollRight": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollRight);
    },

    "call_AutomationNode_ScrollRight": (self: heap.Ref<object>, retPtr: Pointer, callback: heap.Ref<object>): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollRight(A.H.get<object>(callback));
    },
    "try_AutomationNode_ScrollRight": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      callback: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollRight(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollRight1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollRight" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollRight1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollRight);
    },

    "call_AutomationNode_ScrollRight1": (self: heap.Ref<object>, retPtr: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollRight();
    },
    "try_AutomationNode_ScrollRight1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollRight();
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_ScrollToPoint": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "scrollToPoint" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_ScrollToPoint": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).scrollToPoint);
    },

    "call_AutomationNode_ScrollToPoint": (self: heap.Ref<object>, retPtr: Pointer, x: number, y: number): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.scrollToPoint(x, y);
    },
    "try_AutomationNode_ScrollToPoint": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      x: number,
      y: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.scrollToPoint(x, y);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_SetScrollOffset": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "setScrollOffset" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_SetScrollOffset": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).setScrollOffset);
    },

    "call_AutomationNode_SetScrollOffset": (self: heap.Ref<object>, retPtr: Pointer, x: number, y: number): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.setScrollOffset(x, y);
    },
    "try_AutomationNode_SetScrollOffset": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      x: number,
      y: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.setScrollOffset(x, y);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_AddEventListener": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "addEventListener" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_AddEventListener": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).addEventListener);
    },

    "call_AutomationNode_AddEventListener": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventType: number,
      listener: heap.Ref<object>,
      capture: heap.Ref<boolean>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.addEventListener(
        eventType > 0 && eventType <= 119
          ? [
              "accessKeyChanged",
              "activeDescendantChanged",
              "alert",
              "ariaAttributeChangedDeprecated",
              "ariaCurrentChanged",
              "atomicChanged",
              "autoCompleteChanged",
              "autocorrectionOccured",
              "autofillAvailabilityChanged",
              "blur",
              "busyChanged",
              "caretBoundsChanged",
              "checkedStateChanged",
              "checkedStateDescriptionChanged",
              "childrenChanged",
              "classNameChanged",
              "clicked",
              "collapsed",
              "controlsChanged",
              "detailsChanged",
              "describedByChanged",
              "descriptionChanged",
              "documentSelectionChanged",
              "documentTitleChanged",
              "dropeffectChanged",
              "editableTextChanged",
              "enabledChanged",
              "endOfTest",
              "expanded",
              "expandedChanged",
              "flowFromChanged",
              "flowToChanged",
              "focus",
              "focusAfterMenuClose",
              "focusChanged",
              "focusContext",
              "grabbedChanged",
              "haspopupChanged",
              "hide",
              "hierarchicalLevelChanged",
              "hitTestResult",
              "hover",
              "ignoredChanged",
              "imageAnnotationChanged",
              "imageFrameUpdated",
              "invalidStatusChanged",
              "keyShortcutsChanged",
              "labeledByChanged",
              "languageChanged",
              "layoutComplete",
              "layoutInvalidated",
              "liveRegionChanged",
              "liveRegionCreated",
              "liveRegionNodeChanged",
              "liveRelevantChanged",
              "liveStatusChanged",
              "loadComplete",
              "loadStart",
              "locationChanged",
              "mediaStartedPlaying",
              "mediaStoppedPlaying",
              "menuEnd",
              "menuItemSelected",
              "menuListValueChanged",
              "menuPopupEnd",
              "menuPopupStart",
              "menuStart",
              "mouseCanceled",
              "mouseDragged",
              "mouseMoved",
              "mousePressed",
              "mouseReleased",
              "multilineStateChanged",
              "multiselectableStateChanged",
              "nameChanged",
              "objectAttributeChanged",
              "orientationChanged",
              "otherAttributeChanged",
              "parentChanged",
              "placeholderChanged",
              "portalActivated",
              "positionInSetChanged",
              "rangeValueChanged",
              "rangeValueMaxChanged",
              "rangeValueMinChanged",
              "rangeValueStepChanged",
              "readonlyChanged",
              "relatedNodeChanged",
              "requiredStateChanged",
              "roleChanged",
              "rowCollapsed",
              "rowCountChanged",
              "rowExpanded",
              "scrollHorizontalPositionChanged",
              "scrollPositionChanged",
              "scrollVerticalPositionChanged",
              "scrolledToAnchor",
              "selectedChanged",
              "selectedChildrenChanged",
              "selectedValueChanged",
              "selection",
              "selectionAdd",
              "selectionRemove",
              "setSizeChanged",
              "show",
              "sortChanged",
              "stateChanged",
              "subtreeCreated",
              "textAttributeChanged",
              "textSelectionChanged",
              "textChanged",
              "tooltipClosed",
              "tooltipOpened",
              "treeChanged",
              "valueInTextFieldChanged",
              "valueChanged",
              "windowActivated",
              "windowDeactivated",
              "windowVisibilityChanged",
            ][eventType - 1]
          : undefined,
        A.H.get<object>(listener),
        capture === A.H.TRUE
      );
    },
    "try_AutomationNode_AddEventListener": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventType: number,
      listener: heap.Ref<object>,
      capture: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.addEventListener(
          eventType > 0 && eventType <= 119
            ? [
                "accessKeyChanged",
                "activeDescendantChanged",
                "alert",
                "ariaAttributeChangedDeprecated",
                "ariaCurrentChanged",
                "atomicChanged",
                "autoCompleteChanged",
                "autocorrectionOccured",
                "autofillAvailabilityChanged",
                "blur",
                "busyChanged",
                "caretBoundsChanged",
                "checkedStateChanged",
                "checkedStateDescriptionChanged",
                "childrenChanged",
                "classNameChanged",
                "clicked",
                "collapsed",
                "controlsChanged",
                "detailsChanged",
                "describedByChanged",
                "descriptionChanged",
                "documentSelectionChanged",
                "documentTitleChanged",
                "dropeffectChanged",
                "editableTextChanged",
                "enabledChanged",
                "endOfTest",
                "expanded",
                "expandedChanged",
                "flowFromChanged",
                "flowToChanged",
                "focus",
                "focusAfterMenuClose",
                "focusChanged",
                "focusContext",
                "grabbedChanged",
                "haspopupChanged",
                "hide",
                "hierarchicalLevelChanged",
                "hitTestResult",
                "hover",
                "ignoredChanged",
                "imageAnnotationChanged",
                "imageFrameUpdated",
                "invalidStatusChanged",
                "keyShortcutsChanged",
                "labeledByChanged",
                "languageChanged",
                "layoutComplete",
                "layoutInvalidated",
                "liveRegionChanged",
                "liveRegionCreated",
                "liveRegionNodeChanged",
                "liveRelevantChanged",
                "liveStatusChanged",
                "loadComplete",
                "loadStart",
                "locationChanged",
                "mediaStartedPlaying",
                "mediaStoppedPlaying",
                "menuEnd",
                "menuItemSelected",
                "menuListValueChanged",
                "menuPopupEnd",
                "menuPopupStart",
                "menuStart",
                "mouseCanceled",
                "mouseDragged",
                "mouseMoved",
                "mousePressed",
                "mouseReleased",
                "multilineStateChanged",
                "multiselectableStateChanged",
                "nameChanged",
                "objectAttributeChanged",
                "orientationChanged",
                "otherAttributeChanged",
                "parentChanged",
                "placeholderChanged",
                "portalActivated",
                "positionInSetChanged",
                "rangeValueChanged",
                "rangeValueMaxChanged",
                "rangeValueMinChanged",
                "rangeValueStepChanged",
                "readonlyChanged",
                "relatedNodeChanged",
                "requiredStateChanged",
                "roleChanged",
                "rowCollapsed",
                "rowCountChanged",
                "rowExpanded",
                "scrollHorizontalPositionChanged",
                "scrollPositionChanged",
                "scrollVerticalPositionChanged",
                "scrolledToAnchor",
                "selectedChanged",
                "selectedChildrenChanged",
                "selectedValueChanged",
                "selection",
                "selectionAdd",
                "selectionRemove",
                "setSizeChanged",
                "show",
                "sortChanged",
                "stateChanged",
                "subtreeCreated",
                "textAttributeChanged",
                "textSelectionChanged",
                "textChanged",
                "tooltipClosed",
                "tooltipOpened",
                "treeChanged",
                "valueInTextFieldChanged",
                "valueChanged",
                "windowActivated",
                "windowDeactivated",
                "windowVisibilityChanged",
              ][eventType - 1]
            : undefined,
          A.H.get<object>(listener),
          capture === A.H.TRUE
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_RemoveEventListener": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "removeEventListener" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_RemoveEventListener": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).removeEventListener);
    },

    "call_AutomationNode_RemoveEventListener": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      eventType: number,
      listener: heap.Ref<object>,
      capture: heap.Ref<boolean>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.removeEventListener(
        eventType > 0 && eventType <= 119
          ? [
              "accessKeyChanged",
              "activeDescendantChanged",
              "alert",
              "ariaAttributeChangedDeprecated",
              "ariaCurrentChanged",
              "atomicChanged",
              "autoCompleteChanged",
              "autocorrectionOccured",
              "autofillAvailabilityChanged",
              "blur",
              "busyChanged",
              "caretBoundsChanged",
              "checkedStateChanged",
              "checkedStateDescriptionChanged",
              "childrenChanged",
              "classNameChanged",
              "clicked",
              "collapsed",
              "controlsChanged",
              "detailsChanged",
              "describedByChanged",
              "descriptionChanged",
              "documentSelectionChanged",
              "documentTitleChanged",
              "dropeffectChanged",
              "editableTextChanged",
              "enabledChanged",
              "endOfTest",
              "expanded",
              "expandedChanged",
              "flowFromChanged",
              "flowToChanged",
              "focus",
              "focusAfterMenuClose",
              "focusChanged",
              "focusContext",
              "grabbedChanged",
              "haspopupChanged",
              "hide",
              "hierarchicalLevelChanged",
              "hitTestResult",
              "hover",
              "ignoredChanged",
              "imageAnnotationChanged",
              "imageFrameUpdated",
              "invalidStatusChanged",
              "keyShortcutsChanged",
              "labeledByChanged",
              "languageChanged",
              "layoutComplete",
              "layoutInvalidated",
              "liveRegionChanged",
              "liveRegionCreated",
              "liveRegionNodeChanged",
              "liveRelevantChanged",
              "liveStatusChanged",
              "loadComplete",
              "loadStart",
              "locationChanged",
              "mediaStartedPlaying",
              "mediaStoppedPlaying",
              "menuEnd",
              "menuItemSelected",
              "menuListValueChanged",
              "menuPopupEnd",
              "menuPopupStart",
              "menuStart",
              "mouseCanceled",
              "mouseDragged",
              "mouseMoved",
              "mousePressed",
              "mouseReleased",
              "multilineStateChanged",
              "multiselectableStateChanged",
              "nameChanged",
              "objectAttributeChanged",
              "orientationChanged",
              "otherAttributeChanged",
              "parentChanged",
              "placeholderChanged",
              "portalActivated",
              "positionInSetChanged",
              "rangeValueChanged",
              "rangeValueMaxChanged",
              "rangeValueMinChanged",
              "rangeValueStepChanged",
              "readonlyChanged",
              "relatedNodeChanged",
              "requiredStateChanged",
              "roleChanged",
              "rowCollapsed",
              "rowCountChanged",
              "rowExpanded",
              "scrollHorizontalPositionChanged",
              "scrollPositionChanged",
              "scrollVerticalPositionChanged",
              "scrolledToAnchor",
              "selectedChanged",
              "selectedChildrenChanged",
              "selectedValueChanged",
              "selection",
              "selectionAdd",
              "selectionRemove",
              "setSizeChanged",
              "show",
              "sortChanged",
              "stateChanged",
              "subtreeCreated",
              "textAttributeChanged",
              "textSelectionChanged",
              "textChanged",
              "tooltipClosed",
              "tooltipOpened",
              "treeChanged",
              "valueInTextFieldChanged",
              "valueChanged",
              "windowActivated",
              "windowDeactivated",
              "windowVisibilityChanged",
            ][eventType - 1]
          : undefined,
        A.H.get<object>(listener),
        capture === A.H.TRUE
      );
    },
    "try_AutomationNode_RemoveEventListener": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      eventType: number,
      listener: heap.Ref<object>,
      capture: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.removeEventListener(
          eventType > 0 && eventType <= 119
            ? [
                "accessKeyChanged",
                "activeDescendantChanged",
                "alert",
                "ariaAttributeChangedDeprecated",
                "ariaCurrentChanged",
                "atomicChanged",
                "autoCompleteChanged",
                "autocorrectionOccured",
                "autofillAvailabilityChanged",
                "blur",
                "busyChanged",
                "caretBoundsChanged",
                "checkedStateChanged",
                "checkedStateDescriptionChanged",
                "childrenChanged",
                "classNameChanged",
                "clicked",
                "collapsed",
                "controlsChanged",
                "detailsChanged",
                "describedByChanged",
                "descriptionChanged",
                "documentSelectionChanged",
                "documentTitleChanged",
                "dropeffectChanged",
                "editableTextChanged",
                "enabledChanged",
                "endOfTest",
                "expanded",
                "expandedChanged",
                "flowFromChanged",
                "flowToChanged",
                "focus",
                "focusAfterMenuClose",
                "focusChanged",
                "focusContext",
                "grabbedChanged",
                "haspopupChanged",
                "hide",
                "hierarchicalLevelChanged",
                "hitTestResult",
                "hover",
                "ignoredChanged",
                "imageAnnotationChanged",
                "imageFrameUpdated",
                "invalidStatusChanged",
                "keyShortcutsChanged",
                "labeledByChanged",
                "languageChanged",
                "layoutComplete",
                "layoutInvalidated",
                "liveRegionChanged",
                "liveRegionCreated",
                "liveRegionNodeChanged",
                "liveRelevantChanged",
                "liveStatusChanged",
                "loadComplete",
                "loadStart",
                "locationChanged",
                "mediaStartedPlaying",
                "mediaStoppedPlaying",
                "menuEnd",
                "menuItemSelected",
                "menuListValueChanged",
                "menuPopupEnd",
                "menuPopupStart",
                "menuStart",
                "mouseCanceled",
                "mouseDragged",
                "mouseMoved",
                "mousePressed",
                "mouseReleased",
                "multilineStateChanged",
                "multiselectableStateChanged",
                "nameChanged",
                "objectAttributeChanged",
                "orientationChanged",
                "otherAttributeChanged",
                "parentChanged",
                "placeholderChanged",
                "portalActivated",
                "positionInSetChanged",
                "rangeValueChanged",
                "rangeValueMaxChanged",
                "rangeValueMinChanged",
                "rangeValueStepChanged",
                "readonlyChanged",
                "relatedNodeChanged",
                "requiredStateChanged",
                "roleChanged",
                "rowCollapsed",
                "rowCountChanged",
                "rowExpanded",
                "scrollHorizontalPositionChanged",
                "scrollPositionChanged",
                "scrollVerticalPositionChanged",
                "scrolledToAnchor",
                "selectedChanged",
                "selectedChildrenChanged",
                "selectedValueChanged",
                "selection",
                "selectionAdd",
                "selectionRemove",
                "setSizeChanged",
                "show",
                "sortChanged",
                "stateChanged",
                "subtreeCreated",
                "textAttributeChanged",
                "textSelectionChanged",
                "textChanged",
                "tooltipClosed",
                "tooltipOpened",
                "treeChanged",
                "valueInTextFieldChanged",
                "valueChanged",
                "windowActivated",
                "windowDeactivated",
                "windowVisibilityChanged",
              ][eventType - 1]
            : undefined,
          A.H.get<object>(listener),
          capture === A.H.TRUE
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_Find": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "find" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_Find": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).find);
    },

    "call_AutomationNode_Find": (self: heap.Ref<object>, retPtr: Pointer, params: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const params_ffi = {};

      params_ffi["role"] = A.load.Enum(params + 0, [
        "abbr",
        "alert",
        "alertDialog",
        "application",
        "article",
        "audio",
        "banner",
        "blockquote",
        "button",
        "canvas",
        "caption",
        "caret",
        "cell",
        "checkBox",
        "client",
        "code",
        "colorWell",
        "column",
        "columnHeader",
        "comboBoxGrouping",
        "comboBoxMenuButton",
        "comboBoxSelect",
        "comment",
        "complementary",
        "contentDeletion",
        "contentInsertion",
        "contentInfo",
        "date",
        "dateTime",
        "definition",
        "descriptionList",
        "descriptionListDetail",
        "descriptionListTerm",
        "desktop",
        "details",
        "dialog",
        "directory",
        "disclosureTriangle",
        "docAbstract",
        "docAcknowledgments",
        "docAfterword",
        "docAppendix",
        "docBackLink",
        "docBiblioEntry",
        "docBibliography",
        "docBiblioRef",
        "docChapter",
        "docColophon",
        "docConclusion",
        "docCover",
        "docCredit",
        "docCredits",
        "docDedication",
        "docEndnote",
        "docEndnotes",
        "docEpigraph",
        "docEpilogue",
        "docErrata",
        "docExample",
        "docFootnote",
        "docForeword",
        "docGlossary",
        "docGlossRef",
        "docIndex",
        "docIntroduction",
        "docNoteRef",
        "docNotice",
        "docPageBreak",
        "docPageFooter",
        "docPageHeader",
        "docPageList",
        "docPart",
        "docPreface",
        "docPrologue",
        "docPullquote",
        "docQna",
        "docSubtitle",
        "docTip",
        "docToc",
        "document",
        "embeddedObject",
        "emphasis",
        "feed",
        "figcaption",
        "figure",
        "footer",
        "footerAsNonLandmark",
        "form",
        "genericContainer",
        "graphicsDocument",
        "graphicsObject",
        "graphicsSymbol",
        "grid",
        "group",
        "header",
        "headerAsNonLandmark",
        "heading",
        "iframe",
        "iframePresentational",
        "image",
        "imeCandidate",
        "inlineTextBox",
        "inputTime",
        "keyboard",
        "labelText",
        "layoutTable",
        "layoutTableCell",
        "layoutTableRow",
        "legend",
        "lineBreak",
        "link",
        "list",
        "listBox",
        "listBoxOption",
        "listGrid",
        "listItem",
        "listMarker",
        "log",
        "main",
        "mark",
        "marquee",
        "math",
        "mathMLFraction",
        "mathMLIdentifier",
        "mathMLMath",
        "mathMLMultiscripts",
        "mathMLNoneScript",
        "mathMLNumber",
        "mathMLOperator",
        "mathMLOver",
        "mathMLPrescriptDelimiter",
        "mathMLRoot",
        "mathMLRow",
        "mathMLSquareRoot",
        "mathMLStringLiteral",
        "mathMLSub",
        "mathMLSubSup",
        "mathMLSup",
        "mathMLTable",
        "mathMLTableCell",
        "mathMLTableRow",
        "mathMLText",
        "mathMLUnder",
        "mathMLUnderOver",
        "menu",
        "menuBar",
        "menuItem",
        "menuItemCheckBox",
        "menuItemRadio",
        "menuListOption",
        "menuListPopup",
        "meter",
        "navigation",
        "note",
        "pane",
        "paragraph",
        "pdfActionableHighlight",
        "pdfRoot",
        "pluginObject",
        "popUpButton",
        "portal",
        "preDeprecated",
        "progressIndicator",
        "radioButton",
        "radioGroup",
        "region",
        "rootWebArea",
        "row",
        "rowGroup",
        "rowHeader",
        "ruby",
        "rubyAnnotation",
        "scrollBar",
        "scrollView",
        "search",
        "searchBox",
        "section",
        "slider",
        "spinButton",
        "splitter",
        "staticText",
        "status",
        "strong",
        "subscript",
        "suggestion",
        "superscript",
        "svgRoot",
        "switch",
        "tab",
        "tabList",
        "tabPanel",
        "table",
        "tableHeaderContainer",
        "term",
        "textField",
        "textFieldWithComboBox",
        "time",
        "timer",
        "titleBar",
        "toggleButton",
        "toolbar",
        "tooltip",
        "tree",
        "treeGrid",
        "treeItem",
        "unknown",
        "video",
        "webView",
        "window",
      ]);
      params_ffi["state"] = A.load.Ref(params + 4, undefined);
      params_ffi["attributes"] = A.load.Ref(params + 8, undefined);
      const _ret = thiz.find(params_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_AutomationNode_Find": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      params: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const params_ffi = {};

        params_ffi["role"] = A.load.Enum(params + 0, [
          "abbr",
          "alert",
          "alertDialog",
          "application",
          "article",
          "audio",
          "banner",
          "blockquote",
          "button",
          "canvas",
          "caption",
          "caret",
          "cell",
          "checkBox",
          "client",
          "code",
          "colorWell",
          "column",
          "columnHeader",
          "comboBoxGrouping",
          "comboBoxMenuButton",
          "comboBoxSelect",
          "comment",
          "complementary",
          "contentDeletion",
          "contentInsertion",
          "contentInfo",
          "date",
          "dateTime",
          "definition",
          "descriptionList",
          "descriptionListDetail",
          "descriptionListTerm",
          "desktop",
          "details",
          "dialog",
          "directory",
          "disclosureTriangle",
          "docAbstract",
          "docAcknowledgments",
          "docAfterword",
          "docAppendix",
          "docBackLink",
          "docBiblioEntry",
          "docBibliography",
          "docBiblioRef",
          "docChapter",
          "docColophon",
          "docConclusion",
          "docCover",
          "docCredit",
          "docCredits",
          "docDedication",
          "docEndnote",
          "docEndnotes",
          "docEpigraph",
          "docEpilogue",
          "docErrata",
          "docExample",
          "docFootnote",
          "docForeword",
          "docGlossary",
          "docGlossRef",
          "docIndex",
          "docIntroduction",
          "docNoteRef",
          "docNotice",
          "docPageBreak",
          "docPageFooter",
          "docPageHeader",
          "docPageList",
          "docPart",
          "docPreface",
          "docPrologue",
          "docPullquote",
          "docQna",
          "docSubtitle",
          "docTip",
          "docToc",
          "document",
          "embeddedObject",
          "emphasis",
          "feed",
          "figcaption",
          "figure",
          "footer",
          "footerAsNonLandmark",
          "form",
          "genericContainer",
          "graphicsDocument",
          "graphicsObject",
          "graphicsSymbol",
          "grid",
          "group",
          "header",
          "headerAsNonLandmark",
          "heading",
          "iframe",
          "iframePresentational",
          "image",
          "imeCandidate",
          "inlineTextBox",
          "inputTime",
          "keyboard",
          "labelText",
          "layoutTable",
          "layoutTableCell",
          "layoutTableRow",
          "legend",
          "lineBreak",
          "link",
          "list",
          "listBox",
          "listBoxOption",
          "listGrid",
          "listItem",
          "listMarker",
          "log",
          "main",
          "mark",
          "marquee",
          "math",
          "mathMLFraction",
          "mathMLIdentifier",
          "mathMLMath",
          "mathMLMultiscripts",
          "mathMLNoneScript",
          "mathMLNumber",
          "mathMLOperator",
          "mathMLOver",
          "mathMLPrescriptDelimiter",
          "mathMLRoot",
          "mathMLRow",
          "mathMLSquareRoot",
          "mathMLStringLiteral",
          "mathMLSub",
          "mathMLSubSup",
          "mathMLSup",
          "mathMLTable",
          "mathMLTableCell",
          "mathMLTableRow",
          "mathMLText",
          "mathMLUnder",
          "mathMLUnderOver",
          "menu",
          "menuBar",
          "menuItem",
          "menuItemCheckBox",
          "menuItemRadio",
          "menuListOption",
          "menuListPopup",
          "meter",
          "navigation",
          "note",
          "pane",
          "paragraph",
          "pdfActionableHighlight",
          "pdfRoot",
          "pluginObject",
          "popUpButton",
          "portal",
          "preDeprecated",
          "progressIndicator",
          "radioButton",
          "radioGroup",
          "region",
          "rootWebArea",
          "row",
          "rowGroup",
          "rowHeader",
          "ruby",
          "rubyAnnotation",
          "scrollBar",
          "scrollView",
          "search",
          "searchBox",
          "section",
          "slider",
          "spinButton",
          "splitter",
          "staticText",
          "status",
          "strong",
          "subscript",
          "suggestion",
          "superscript",
          "svgRoot",
          "switch",
          "tab",
          "tabList",
          "tabPanel",
          "table",
          "tableHeaderContainer",
          "term",
          "textField",
          "textFieldWithComboBox",
          "time",
          "timer",
          "titleBar",
          "toggleButton",
          "toolbar",
          "tooltip",
          "tree",
          "treeGrid",
          "treeItem",
          "unknown",
          "video",
          "webView",
          "window",
        ]);
        params_ffi["state"] = A.load.Ref(params + 4, undefined);
        params_ffi["attributes"] = A.load.Ref(params + 8, undefined);
        const _ret = thiz.find(params_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_FindAll": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "findAll" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_FindAll": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).findAll);
    },

    "call_AutomationNode_FindAll": (self: heap.Ref<object>, retPtr: Pointer, params: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const params_ffi = {};

      params_ffi["role"] = A.load.Enum(params + 0, [
        "abbr",
        "alert",
        "alertDialog",
        "application",
        "article",
        "audio",
        "banner",
        "blockquote",
        "button",
        "canvas",
        "caption",
        "caret",
        "cell",
        "checkBox",
        "client",
        "code",
        "colorWell",
        "column",
        "columnHeader",
        "comboBoxGrouping",
        "comboBoxMenuButton",
        "comboBoxSelect",
        "comment",
        "complementary",
        "contentDeletion",
        "contentInsertion",
        "contentInfo",
        "date",
        "dateTime",
        "definition",
        "descriptionList",
        "descriptionListDetail",
        "descriptionListTerm",
        "desktop",
        "details",
        "dialog",
        "directory",
        "disclosureTriangle",
        "docAbstract",
        "docAcknowledgments",
        "docAfterword",
        "docAppendix",
        "docBackLink",
        "docBiblioEntry",
        "docBibliography",
        "docBiblioRef",
        "docChapter",
        "docColophon",
        "docConclusion",
        "docCover",
        "docCredit",
        "docCredits",
        "docDedication",
        "docEndnote",
        "docEndnotes",
        "docEpigraph",
        "docEpilogue",
        "docErrata",
        "docExample",
        "docFootnote",
        "docForeword",
        "docGlossary",
        "docGlossRef",
        "docIndex",
        "docIntroduction",
        "docNoteRef",
        "docNotice",
        "docPageBreak",
        "docPageFooter",
        "docPageHeader",
        "docPageList",
        "docPart",
        "docPreface",
        "docPrologue",
        "docPullquote",
        "docQna",
        "docSubtitle",
        "docTip",
        "docToc",
        "document",
        "embeddedObject",
        "emphasis",
        "feed",
        "figcaption",
        "figure",
        "footer",
        "footerAsNonLandmark",
        "form",
        "genericContainer",
        "graphicsDocument",
        "graphicsObject",
        "graphicsSymbol",
        "grid",
        "group",
        "header",
        "headerAsNonLandmark",
        "heading",
        "iframe",
        "iframePresentational",
        "image",
        "imeCandidate",
        "inlineTextBox",
        "inputTime",
        "keyboard",
        "labelText",
        "layoutTable",
        "layoutTableCell",
        "layoutTableRow",
        "legend",
        "lineBreak",
        "link",
        "list",
        "listBox",
        "listBoxOption",
        "listGrid",
        "listItem",
        "listMarker",
        "log",
        "main",
        "mark",
        "marquee",
        "math",
        "mathMLFraction",
        "mathMLIdentifier",
        "mathMLMath",
        "mathMLMultiscripts",
        "mathMLNoneScript",
        "mathMLNumber",
        "mathMLOperator",
        "mathMLOver",
        "mathMLPrescriptDelimiter",
        "mathMLRoot",
        "mathMLRow",
        "mathMLSquareRoot",
        "mathMLStringLiteral",
        "mathMLSub",
        "mathMLSubSup",
        "mathMLSup",
        "mathMLTable",
        "mathMLTableCell",
        "mathMLTableRow",
        "mathMLText",
        "mathMLUnder",
        "mathMLUnderOver",
        "menu",
        "menuBar",
        "menuItem",
        "menuItemCheckBox",
        "menuItemRadio",
        "menuListOption",
        "menuListPopup",
        "meter",
        "navigation",
        "note",
        "pane",
        "paragraph",
        "pdfActionableHighlight",
        "pdfRoot",
        "pluginObject",
        "popUpButton",
        "portal",
        "preDeprecated",
        "progressIndicator",
        "radioButton",
        "radioGroup",
        "region",
        "rootWebArea",
        "row",
        "rowGroup",
        "rowHeader",
        "ruby",
        "rubyAnnotation",
        "scrollBar",
        "scrollView",
        "search",
        "searchBox",
        "section",
        "slider",
        "spinButton",
        "splitter",
        "staticText",
        "status",
        "strong",
        "subscript",
        "suggestion",
        "superscript",
        "svgRoot",
        "switch",
        "tab",
        "tabList",
        "tabPanel",
        "table",
        "tableHeaderContainer",
        "term",
        "textField",
        "textFieldWithComboBox",
        "time",
        "timer",
        "titleBar",
        "toggleButton",
        "toolbar",
        "tooltip",
        "tree",
        "treeGrid",
        "treeItem",
        "unknown",
        "video",
        "webView",
        "window",
      ]);
      params_ffi["state"] = A.load.Ref(params + 4, undefined);
      params_ffi["attributes"] = A.load.Ref(params + 8, undefined);
      const _ret = thiz.findAll(params_ffi);
      A.store.Ref(retPtr, _ret);
    },
    "try_AutomationNode_FindAll": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      params: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const params_ffi = {};

        params_ffi["role"] = A.load.Enum(params + 0, [
          "abbr",
          "alert",
          "alertDialog",
          "application",
          "article",
          "audio",
          "banner",
          "blockquote",
          "button",
          "canvas",
          "caption",
          "caret",
          "cell",
          "checkBox",
          "client",
          "code",
          "colorWell",
          "column",
          "columnHeader",
          "comboBoxGrouping",
          "comboBoxMenuButton",
          "comboBoxSelect",
          "comment",
          "complementary",
          "contentDeletion",
          "contentInsertion",
          "contentInfo",
          "date",
          "dateTime",
          "definition",
          "descriptionList",
          "descriptionListDetail",
          "descriptionListTerm",
          "desktop",
          "details",
          "dialog",
          "directory",
          "disclosureTriangle",
          "docAbstract",
          "docAcknowledgments",
          "docAfterword",
          "docAppendix",
          "docBackLink",
          "docBiblioEntry",
          "docBibliography",
          "docBiblioRef",
          "docChapter",
          "docColophon",
          "docConclusion",
          "docCover",
          "docCredit",
          "docCredits",
          "docDedication",
          "docEndnote",
          "docEndnotes",
          "docEpigraph",
          "docEpilogue",
          "docErrata",
          "docExample",
          "docFootnote",
          "docForeword",
          "docGlossary",
          "docGlossRef",
          "docIndex",
          "docIntroduction",
          "docNoteRef",
          "docNotice",
          "docPageBreak",
          "docPageFooter",
          "docPageHeader",
          "docPageList",
          "docPart",
          "docPreface",
          "docPrologue",
          "docPullquote",
          "docQna",
          "docSubtitle",
          "docTip",
          "docToc",
          "document",
          "embeddedObject",
          "emphasis",
          "feed",
          "figcaption",
          "figure",
          "footer",
          "footerAsNonLandmark",
          "form",
          "genericContainer",
          "graphicsDocument",
          "graphicsObject",
          "graphicsSymbol",
          "grid",
          "group",
          "header",
          "headerAsNonLandmark",
          "heading",
          "iframe",
          "iframePresentational",
          "image",
          "imeCandidate",
          "inlineTextBox",
          "inputTime",
          "keyboard",
          "labelText",
          "layoutTable",
          "layoutTableCell",
          "layoutTableRow",
          "legend",
          "lineBreak",
          "link",
          "list",
          "listBox",
          "listBoxOption",
          "listGrid",
          "listItem",
          "listMarker",
          "log",
          "main",
          "mark",
          "marquee",
          "math",
          "mathMLFraction",
          "mathMLIdentifier",
          "mathMLMath",
          "mathMLMultiscripts",
          "mathMLNoneScript",
          "mathMLNumber",
          "mathMLOperator",
          "mathMLOver",
          "mathMLPrescriptDelimiter",
          "mathMLRoot",
          "mathMLRow",
          "mathMLSquareRoot",
          "mathMLStringLiteral",
          "mathMLSub",
          "mathMLSubSup",
          "mathMLSup",
          "mathMLTable",
          "mathMLTableCell",
          "mathMLTableRow",
          "mathMLText",
          "mathMLUnder",
          "mathMLUnderOver",
          "menu",
          "menuBar",
          "menuItem",
          "menuItemCheckBox",
          "menuItemRadio",
          "menuListOption",
          "menuListPopup",
          "meter",
          "navigation",
          "note",
          "pane",
          "paragraph",
          "pdfActionableHighlight",
          "pdfRoot",
          "pluginObject",
          "popUpButton",
          "portal",
          "preDeprecated",
          "progressIndicator",
          "radioButton",
          "radioGroup",
          "region",
          "rootWebArea",
          "row",
          "rowGroup",
          "rowHeader",
          "ruby",
          "rubyAnnotation",
          "scrollBar",
          "scrollView",
          "search",
          "searchBox",
          "section",
          "slider",
          "spinButton",
          "splitter",
          "staticText",
          "status",
          "strong",
          "subscript",
          "suggestion",
          "superscript",
          "svgRoot",
          "switch",
          "tab",
          "tabList",
          "tabPanel",
          "table",
          "tableHeaderContainer",
          "term",
          "textField",
          "textFieldWithComboBox",
          "time",
          "timer",
          "titleBar",
          "toggleButton",
          "toolbar",
          "tooltip",
          "tree",
          "treeGrid",
          "treeItem",
          "unknown",
          "video",
          "webView",
          "window",
        ]);
        params_ffi["state"] = A.load.Ref(params + 4, undefined);
        params_ffi["attributes"] = A.load.Ref(params + 8, undefined);
        const _ret = thiz.findAll(params_ffi);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_Matches": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "matches" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_Matches": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).matches);
    },

    "call_AutomationNode_Matches": (self: heap.Ref<object>, retPtr: Pointer, params: Pointer): void => {
      const thiz = A.H.get<any>(self);

      const params_ffi = {};

      params_ffi["role"] = A.load.Enum(params + 0, [
        "abbr",
        "alert",
        "alertDialog",
        "application",
        "article",
        "audio",
        "banner",
        "blockquote",
        "button",
        "canvas",
        "caption",
        "caret",
        "cell",
        "checkBox",
        "client",
        "code",
        "colorWell",
        "column",
        "columnHeader",
        "comboBoxGrouping",
        "comboBoxMenuButton",
        "comboBoxSelect",
        "comment",
        "complementary",
        "contentDeletion",
        "contentInsertion",
        "contentInfo",
        "date",
        "dateTime",
        "definition",
        "descriptionList",
        "descriptionListDetail",
        "descriptionListTerm",
        "desktop",
        "details",
        "dialog",
        "directory",
        "disclosureTriangle",
        "docAbstract",
        "docAcknowledgments",
        "docAfterword",
        "docAppendix",
        "docBackLink",
        "docBiblioEntry",
        "docBibliography",
        "docBiblioRef",
        "docChapter",
        "docColophon",
        "docConclusion",
        "docCover",
        "docCredit",
        "docCredits",
        "docDedication",
        "docEndnote",
        "docEndnotes",
        "docEpigraph",
        "docEpilogue",
        "docErrata",
        "docExample",
        "docFootnote",
        "docForeword",
        "docGlossary",
        "docGlossRef",
        "docIndex",
        "docIntroduction",
        "docNoteRef",
        "docNotice",
        "docPageBreak",
        "docPageFooter",
        "docPageHeader",
        "docPageList",
        "docPart",
        "docPreface",
        "docPrologue",
        "docPullquote",
        "docQna",
        "docSubtitle",
        "docTip",
        "docToc",
        "document",
        "embeddedObject",
        "emphasis",
        "feed",
        "figcaption",
        "figure",
        "footer",
        "footerAsNonLandmark",
        "form",
        "genericContainer",
        "graphicsDocument",
        "graphicsObject",
        "graphicsSymbol",
        "grid",
        "group",
        "header",
        "headerAsNonLandmark",
        "heading",
        "iframe",
        "iframePresentational",
        "image",
        "imeCandidate",
        "inlineTextBox",
        "inputTime",
        "keyboard",
        "labelText",
        "layoutTable",
        "layoutTableCell",
        "layoutTableRow",
        "legend",
        "lineBreak",
        "link",
        "list",
        "listBox",
        "listBoxOption",
        "listGrid",
        "listItem",
        "listMarker",
        "log",
        "main",
        "mark",
        "marquee",
        "math",
        "mathMLFraction",
        "mathMLIdentifier",
        "mathMLMath",
        "mathMLMultiscripts",
        "mathMLNoneScript",
        "mathMLNumber",
        "mathMLOperator",
        "mathMLOver",
        "mathMLPrescriptDelimiter",
        "mathMLRoot",
        "mathMLRow",
        "mathMLSquareRoot",
        "mathMLStringLiteral",
        "mathMLSub",
        "mathMLSubSup",
        "mathMLSup",
        "mathMLTable",
        "mathMLTableCell",
        "mathMLTableRow",
        "mathMLText",
        "mathMLUnder",
        "mathMLUnderOver",
        "menu",
        "menuBar",
        "menuItem",
        "menuItemCheckBox",
        "menuItemRadio",
        "menuListOption",
        "menuListPopup",
        "meter",
        "navigation",
        "note",
        "pane",
        "paragraph",
        "pdfActionableHighlight",
        "pdfRoot",
        "pluginObject",
        "popUpButton",
        "portal",
        "preDeprecated",
        "progressIndicator",
        "radioButton",
        "radioGroup",
        "region",
        "rootWebArea",
        "row",
        "rowGroup",
        "rowHeader",
        "ruby",
        "rubyAnnotation",
        "scrollBar",
        "scrollView",
        "search",
        "searchBox",
        "section",
        "slider",
        "spinButton",
        "splitter",
        "staticText",
        "status",
        "strong",
        "subscript",
        "suggestion",
        "superscript",
        "svgRoot",
        "switch",
        "tab",
        "tabList",
        "tabPanel",
        "table",
        "tableHeaderContainer",
        "term",
        "textField",
        "textFieldWithComboBox",
        "time",
        "timer",
        "titleBar",
        "toggleButton",
        "toolbar",
        "tooltip",
        "tree",
        "treeGrid",
        "treeItem",
        "unknown",
        "video",
        "webView",
        "window",
      ]);
      params_ffi["state"] = A.load.Ref(params + 4, undefined);
      params_ffi["attributes"] = A.load.Ref(params + 8, undefined);
      const _ret = thiz.matches(params_ffi);
      A.store.Bool(retPtr, _ret);
    },
    "try_AutomationNode_Matches": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      params: Pointer
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const params_ffi = {};

        params_ffi["role"] = A.load.Enum(params + 0, [
          "abbr",
          "alert",
          "alertDialog",
          "application",
          "article",
          "audio",
          "banner",
          "blockquote",
          "button",
          "canvas",
          "caption",
          "caret",
          "cell",
          "checkBox",
          "client",
          "code",
          "colorWell",
          "column",
          "columnHeader",
          "comboBoxGrouping",
          "comboBoxMenuButton",
          "comboBoxSelect",
          "comment",
          "complementary",
          "contentDeletion",
          "contentInsertion",
          "contentInfo",
          "date",
          "dateTime",
          "definition",
          "descriptionList",
          "descriptionListDetail",
          "descriptionListTerm",
          "desktop",
          "details",
          "dialog",
          "directory",
          "disclosureTriangle",
          "docAbstract",
          "docAcknowledgments",
          "docAfterword",
          "docAppendix",
          "docBackLink",
          "docBiblioEntry",
          "docBibliography",
          "docBiblioRef",
          "docChapter",
          "docColophon",
          "docConclusion",
          "docCover",
          "docCredit",
          "docCredits",
          "docDedication",
          "docEndnote",
          "docEndnotes",
          "docEpigraph",
          "docEpilogue",
          "docErrata",
          "docExample",
          "docFootnote",
          "docForeword",
          "docGlossary",
          "docGlossRef",
          "docIndex",
          "docIntroduction",
          "docNoteRef",
          "docNotice",
          "docPageBreak",
          "docPageFooter",
          "docPageHeader",
          "docPageList",
          "docPart",
          "docPreface",
          "docPrologue",
          "docPullquote",
          "docQna",
          "docSubtitle",
          "docTip",
          "docToc",
          "document",
          "embeddedObject",
          "emphasis",
          "feed",
          "figcaption",
          "figure",
          "footer",
          "footerAsNonLandmark",
          "form",
          "genericContainer",
          "graphicsDocument",
          "graphicsObject",
          "graphicsSymbol",
          "grid",
          "group",
          "header",
          "headerAsNonLandmark",
          "heading",
          "iframe",
          "iframePresentational",
          "image",
          "imeCandidate",
          "inlineTextBox",
          "inputTime",
          "keyboard",
          "labelText",
          "layoutTable",
          "layoutTableCell",
          "layoutTableRow",
          "legend",
          "lineBreak",
          "link",
          "list",
          "listBox",
          "listBoxOption",
          "listGrid",
          "listItem",
          "listMarker",
          "log",
          "main",
          "mark",
          "marquee",
          "math",
          "mathMLFraction",
          "mathMLIdentifier",
          "mathMLMath",
          "mathMLMultiscripts",
          "mathMLNoneScript",
          "mathMLNumber",
          "mathMLOperator",
          "mathMLOver",
          "mathMLPrescriptDelimiter",
          "mathMLRoot",
          "mathMLRow",
          "mathMLSquareRoot",
          "mathMLStringLiteral",
          "mathMLSub",
          "mathMLSubSup",
          "mathMLSup",
          "mathMLTable",
          "mathMLTableCell",
          "mathMLTableRow",
          "mathMLText",
          "mathMLUnder",
          "mathMLUnderOver",
          "menu",
          "menuBar",
          "menuItem",
          "menuItemCheckBox",
          "menuItemRadio",
          "menuListOption",
          "menuListPopup",
          "meter",
          "navigation",
          "note",
          "pane",
          "paragraph",
          "pdfActionableHighlight",
          "pdfRoot",
          "pluginObject",
          "popUpButton",
          "portal",
          "preDeprecated",
          "progressIndicator",
          "radioButton",
          "radioGroup",
          "region",
          "rootWebArea",
          "row",
          "rowGroup",
          "rowHeader",
          "ruby",
          "rubyAnnotation",
          "scrollBar",
          "scrollView",
          "search",
          "searchBox",
          "section",
          "slider",
          "spinButton",
          "splitter",
          "staticText",
          "status",
          "strong",
          "subscript",
          "suggestion",
          "superscript",
          "svgRoot",
          "switch",
          "tab",
          "tabList",
          "tabPanel",
          "table",
          "tableHeaderContainer",
          "term",
          "textField",
          "textFieldWithComboBox",
          "time",
          "timer",
          "titleBar",
          "toggleButton",
          "toolbar",
          "tooltip",
          "tree",
          "treeGrid",
          "treeItem",
          "unknown",
          "video",
          "webView",
          "window",
        ]);
        params_ffi["state"] = A.load.Ref(params + 4, undefined);
        params_ffi["attributes"] = A.load.Ref(params + 8, undefined);
        const _ret = thiz.matches(params_ffi);
        A.store.Bool(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_GetNextTextMatch": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "getNextTextMatch" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_GetNextTextMatch": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).getNextTextMatch);
    },

    "call_AutomationNode_GetNextTextMatch": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      searchStr: heap.Ref<object>,
      backward: heap.Ref<boolean>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.getNextTextMatch(A.H.get<object>(searchStr), backward === A.H.TRUE);
      A.store.Ref(retPtr, _ret);
    },
    "try_AutomationNode_GetNextTextMatch": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      searchStr: heap.Ref<object>,
      backward: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.getNextTextMatch(A.H.get<object>(searchStr), backward === A.H.TRUE);
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_CreatePosition": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "createPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_CreatePosition": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).createPosition);
    },

    "call_AutomationNode_CreatePosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      type: number,
      offset: number,
      isUpstream: heap.Ref<boolean>
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.createPosition(
        type > 0 && type <= 3 ? ["null", "text", "tree"][type - 1] : undefined,
        offset,
        isUpstream === A.H.TRUE
      );
      A.store.Ref(retPtr, _ret);
    },
    "try_AutomationNode_CreatePosition": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      type: number,
      offset: number,
      isUpstream: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.createPosition(
          type > 0 && type <= 3 ? ["null", "text", "tree"][type - 1] : undefined,
          offset,
          isUpstream === A.H.TRUE
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_AutomationNode_CreatePosition1": (self: heap.Ref<any>): heap.Ref<boolean> => {
      const obj = A.H.get<object>(self);
      if (obj && "createPosition" in obj) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AutomationNode_CreatePosition1": (self: heap.Ref<any>, fn: Pointer): void => {
      A.store.Ref(fn, A.H.get<any>(self).createPosition);
    },

    "call_AutomationNode_CreatePosition1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      type: number,
      offset: number
    ): void => {
      const thiz = A.H.get<any>(self);

      const _ret = thiz.createPosition(type > 0 && type <= 3 ? ["null", "text", "tree"][type - 1] : undefined, offset);
      A.store.Ref(retPtr, _ret);
    },
    "try_AutomationNode_CreatePosition1": (
      self: heap.Ref<object>,
      retPtr: Pointer,
      errPtr: Pointer,
      type: number,
      offset: number
    ): heap.Ref<boolean> => {
      try {
        const thiz = A.H.get<any>(self);

        const _ret = thiz.createPosition(
          type > 0 && type <= 3 ? ["null", "text", "tree"][type - 1] : undefined,
          offset
        );
        A.store.Ref(retPtr, _ret);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "get_AutomationNode_Root": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "root" in thiz) {
        const val = thiz.root;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Root": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "root", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_IsRootNode": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "isRootNode" in thiz) {
        const val = thiz.isRootNode;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_IsRootNode": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "isRootNode", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Role": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "role" in thiz) {
        const val = thiz.role;
        A.store.Enum(
          retPtr,
          [
            "abbr",
            "alert",
            "alertDialog",
            "application",
            "article",
            "audio",
            "banner",
            "blockquote",
            "button",
            "canvas",
            "caption",
            "caret",
            "cell",
            "checkBox",
            "client",
            "code",
            "colorWell",
            "column",
            "columnHeader",
            "comboBoxGrouping",
            "comboBoxMenuButton",
            "comboBoxSelect",
            "comment",
            "complementary",
            "contentDeletion",
            "contentInsertion",
            "contentInfo",
            "date",
            "dateTime",
            "definition",
            "descriptionList",
            "descriptionListDetail",
            "descriptionListTerm",
            "desktop",
            "details",
            "dialog",
            "directory",
            "disclosureTriangle",
            "docAbstract",
            "docAcknowledgments",
            "docAfterword",
            "docAppendix",
            "docBackLink",
            "docBiblioEntry",
            "docBibliography",
            "docBiblioRef",
            "docChapter",
            "docColophon",
            "docConclusion",
            "docCover",
            "docCredit",
            "docCredits",
            "docDedication",
            "docEndnote",
            "docEndnotes",
            "docEpigraph",
            "docEpilogue",
            "docErrata",
            "docExample",
            "docFootnote",
            "docForeword",
            "docGlossary",
            "docGlossRef",
            "docIndex",
            "docIntroduction",
            "docNoteRef",
            "docNotice",
            "docPageBreak",
            "docPageFooter",
            "docPageHeader",
            "docPageList",
            "docPart",
            "docPreface",
            "docPrologue",
            "docPullquote",
            "docQna",
            "docSubtitle",
            "docTip",
            "docToc",
            "document",
            "embeddedObject",
            "emphasis",
            "feed",
            "figcaption",
            "figure",
            "footer",
            "footerAsNonLandmark",
            "form",
            "genericContainer",
            "graphicsDocument",
            "graphicsObject",
            "graphicsSymbol",
            "grid",
            "group",
            "header",
            "headerAsNonLandmark",
            "heading",
            "iframe",
            "iframePresentational",
            "image",
            "imeCandidate",
            "inlineTextBox",
            "inputTime",
            "keyboard",
            "labelText",
            "layoutTable",
            "layoutTableCell",
            "layoutTableRow",
            "legend",
            "lineBreak",
            "link",
            "list",
            "listBox",
            "listBoxOption",
            "listGrid",
            "listItem",
            "listMarker",
            "log",
            "main",
            "mark",
            "marquee",
            "math",
            "mathMLFraction",
            "mathMLIdentifier",
            "mathMLMath",
            "mathMLMultiscripts",
            "mathMLNoneScript",
            "mathMLNumber",
            "mathMLOperator",
            "mathMLOver",
            "mathMLPrescriptDelimiter",
            "mathMLRoot",
            "mathMLRow",
            "mathMLSquareRoot",
            "mathMLStringLiteral",
            "mathMLSub",
            "mathMLSubSup",
            "mathMLSup",
            "mathMLTable",
            "mathMLTableCell",
            "mathMLTableRow",
            "mathMLText",
            "mathMLUnder",
            "mathMLUnderOver",
            "menu",
            "menuBar",
            "menuItem",
            "menuItemCheckBox",
            "menuItemRadio",
            "menuListOption",
            "menuListPopup",
            "meter",
            "navigation",
            "note",
            "pane",
            "paragraph",
            "pdfActionableHighlight",
            "pdfRoot",
            "pluginObject",
            "popUpButton",
            "portal",
            "preDeprecated",
            "progressIndicator",
            "radioButton",
            "radioGroup",
            "region",
            "rootWebArea",
            "row",
            "rowGroup",
            "rowHeader",
            "ruby",
            "rubyAnnotation",
            "scrollBar",
            "scrollView",
            "search",
            "searchBox",
            "section",
            "slider",
            "spinButton",
            "splitter",
            "staticText",
            "status",
            "strong",
            "subscript",
            "suggestion",
            "superscript",
            "svgRoot",
            "switch",
            "tab",
            "tabList",
            "tabPanel",
            "table",
            "tableHeaderContainer",
            "term",
            "textField",
            "textFieldWithComboBox",
            "time",
            "timer",
            "titleBar",
            "toggleButton",
            "toolbar",
            "tooltip",
            "tree",
            "treeGrid",
            "treeItem",
            "unknown",
            "video",
            "webView",
            "window",
          ].indexOf(val as string)
        );
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Role": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(
        thiz,
        "role",
        val > 0 && val <= 209
          ? [
              "abbr",
              "alert",
              "alertDialog",
              "application",
              "article",
              "audio",
              "banner",
              "blockquote",
              "button",
              "canvas",
              "caption",
              "caret",
              "cell",
              "checkBox",
              "client",
              "code",
              "colorWell",
              "column",
              "columnHeader",
              "comboBoxGrouping",
              "comboBoxMenuButton",
              "comboBoxSelect",
              "comment",
              "complementary",
              "contentDeletion",
              "contentInsertion",
              "contentInfo",
              "date",
              "dateTime",
              "definition",
              "descriptionList",
              "descriptionListDetail",
              "descriptionListTerm",
              "desktop",
              "details",
              "dialog",
              "directory",
              "disclosureTriangle",
              "docAbstract",
              "docAcknowledgments",
              "docAfterword",
              "docAppendix",
              "docBackLink",
              "docBiblioEntry",
              "docBibliography",
              "docBiblioRef",
              "docChapter",
              "docColophon",
              "docConclusion",
              "docCover",
              "docCredit",
              "docCredits",
              "docDedication",
              "docEndnote",
              "docEndnotes",
              "docEpigraph",
              "docEpilogue",
              "docErrata",
              "docExample",
              "docFootnote",
              "docForeword",
              "docGlossary",
              "docGlossRef",
              "docIndex",
              "docIntroduction",
              "docNoteRef",
              "docNotice",
              "docPageBreak",
              "docPageFooter",
              "docPageHeader",
              "docPageList",
              "docPart",
              "docPreface",
              "docPrologue",
              "docPullquote",
              "docQna",
              "docSubtitle",
              "docTip",
              "docToc",
              "document",
              "embeddedObject",
              "emphasis",
              "feed",
              "figcaption",
              "figure",
              "footer",
              "footerAsNonLandmark",
              "form",
              "genericContainer",
              "graphicsDocument",
              "graphicsObject",
              "graphicsSymbol",
              "grid",
              "group",
              "header",
              "headerAsNonLandmark",
              "heading",
              "iframe",
              "iframePresentational",
              "image",
              "imeCandidate",
              "inlineTextBox",
              "inputTime",
              "keyboard",
              "labelText",
              "layoutTable",
              "layoutTableCell",
              "layoutTableRow",
              "legend",
              "lineBreak",
              "link",
              "list",
              "listBox",
              "listBoxOption",
              "listGrid",
              "listItem",
              "listMarker",
              "log",
              "main",
              "mark",
              "marquee",
              "math",
              "mathMLFraction",
              "mathMLIdentifier",
              "mathMLMath",
              "mathMLMultiscripts",
              "mathMLNoneScript",
              "mathMLNumber",
              "mathMLOperator",
              "mathMLOver",
              "mathMLPrescriptDelimiter",
              "mathMLRoot",
              "mathMLRow",
              "mathMLSquareRoot",
              "mathMLStringLiteral",
              "mathMLSub",
              "mathMLSubSup",
              "mathMLSup",
              "mathMLTable",
              "mathMLTableCell",
              "mathMLTableRow",
              "mathMLText",
              "mathMLUnder",
              "mathMLUnderOver",
              "menu",
              "menuBar",
              "menuItem",
              "menuItemCheckBox",
              "menuItemRadio",
              "menuListOption",
              "menuListPopup",
              "meter",
              "navigation",
              "note",
              "pane",
              "paragraph",
              "pdfActionableHighlight",
              "pdfRoot",
              "pluginObject",
              "popUpButton",
              "portal",
              "preDeprecated",
              "progressIndicator",
              "radioButton",
              "radioGroup",
              "region",
              "rootWebArea",
              "row",
              "rowGroup",
              "rowHeader",
              "ruby",
              "rubyAnnotation",
              "scrollBar",
              "scrollView",
              "search",
              "searchBox",
              "section",
              "slider",
              "spinButton",
              "splitter",
              "staticText",
              "status",
              "strong",
              "subscript",
              "suggestion",
              "superscript",
              "svgRoot",
              "switch",
              "tab",
              "tabList",
              "tabPanel",
              "table",
              "tableHeaderContainer",
              "term",
              "textField",
              "textFieldWithComboBox",
              "time",
              "timer",
              "titleBar",
              "toggleButton",
              "toolbar",
              "tooltip",
              "tree",
              "treeGrid",
              "treeItem",
              "unknown",
              "video",
              "webView",
              "window",
            ][val - 1]
          : undefined,
        thiz
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationNode_State": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "state" in thiz) {
        const val = thiz.state;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_State": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "state", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Location": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "location" in thiz) {
        const val = thiz.location;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 20, false);
          A.store.Bool(retPtr + 16, false);
          A.store.Int32(retPtr + 0, 0);
          A.store.Bool(retPtr + 17, false);
          A.store.Int32(retPtr + 4, 0);
          A.store.Bool(retPtr + 18, false);
          A.store.Int32(retPtr + 8, 0);
          A.store.Bool(retPtr + 19, false);
          A.store.Int32(retPtr + 12, 0);
        } else {
          A.store.Bool(retPtr + 20, true);
          A.store.Bool(retPtr + 16, "left" in val ? true : false);
          A.store.Int32(retPtr + 0, val["left"] === undefined ? 0 : (val["left"] as number));
          A.store.Bool(retPtr + 17, "top" in val ? true : false);
          A.store.Int32(retPtr + 4, val["top"] === undefined ? 0 : (val["top"] as number));
          A.store.Bool(retPtr + 18, "width" in val ? true : false);
          A.store.Int32(retPtr + 8, val["width"] === undefined ? 0 : (val["width"] as number));
          A.store.Bool(retPtr + 19, "height" in val ? true : false);
          A.store.Int32(retPtr + 12, val["height"] === undefined ? 0 : (val["height"] as number));
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Location": (self: heap.Ref<object>, val: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      const val_ffi = {};

      if (A.load.Bool(val + 16)) {
        val_ffi["left"] = A.load.Int32(val + 0);
      }
      if (A.load.Bool(val + 17)) {
        val_ffi["top"] = A.load.Int32(val + 4);
      }
      if (A.load.Bool(val + 18)) {
        val_ffi["width"] = A.load.Int32(val + 8);
      }
      if (A.load.Bool(val + 19)) {
        val_ffi["height"] = A.load.Int32(val + 12);
      }
      return Reflect.set(thiz, "location", val_ffi, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_UnclippedLocation": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "unclippedLocation" in thiz) {
        const val = thiz.unclippedLocation;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 20, false);
          A.store.Bool(retPtr + 16, false);
          A.store.Int32(retPtr + 0, 0);
          A.store.Bool(retPtr + 17, false);
          A.store.Int32(retPtr + 4, 0);
          A.store.Bool(retPtr + 18, false);
          A.store.Int32(retPtr + 8, 0);
          A.store.Bool(retPtr + 19, false);
          A.store.Int32(retPtr + 12, 0);
        } else {
          A.store.Bool(retPtr + 20, true);
          A.store.Bool(retPtr + 16, "left" in val ? true : false);
          A.store.Int32(retPtr + 0, val["left"] === undefined ? 0 : (val["left"] as number));
          A.store.Bool(retPtr + 17, "top" in val ? true : false);
          A.store.Int32(retPtr + 4, val["top"] === undefined ? 0 : (val["top"] as number));
          A.store.Bool(retPtr + 18, "width" in val ? true : false);
          A.store.Int32(retPtr + 8, val["width"] === undefined ? 0 : (val["width"] as number));
          A.store.Bool(retPtr + 19, "height" in val ? true : false);
          A.store.Int32(retPtr + 12, val["height"] === undefined ? 0 : (val["height"] as number));
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_UnclippedLocation": (self: heap.Ref<object>, val: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      const val_ffi = {};

      if (A.load.Bool(val + 16)) {
        val_ffi["left"] = A.load.Int32(val + 0);
      }
      if (A.load.Bool(val + 17)) {
        val_ffi["top"] = A.load.Int32(val + 4);
      }
      if (A.load.Bool(val + 18)) {
        val_ffi["width"] = A.load.Int32(val + 8);
      }
      if (A.load.Bool(val + 19)) {
        val_ffi["height"] = A.load.Int32(val + 12);
      }
      return Reflect.set(thiz, "unclippedLocation", val_ffi, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Description": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "description" in thiz) {
        const val = thiz.description;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Description": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "description", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_CheckedStateDescription": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "checkedStateDescription" in thiz) {
        const val = thiz.checkedStateDescription;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_CheckedStateDescription": (
      self: heap.Ref<object>,
      val: heap.Ref<object>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "checkedStateDescription", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Placeholder": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "placeholder" in thiz) {
        const val = thiz.placeholder;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Placeholder": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "placeholder", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_RoleDescription": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "roleDescription" in thiz) {
        const val = thiz.roleDescription;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_RoleDescription": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "roleDescription", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Name": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "name" in thiz) {
        const val = thiz.name;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Name": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "name", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DoDefaultLabel": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "doDefaultLabel" in thiz) {
        const val = thiz.doDefaultLabel;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DoDefaultLabel": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "doDefaultLabel", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LongClickLabel": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "longClickLabel" in thiz) {
        const val = thiz.longClickLabel;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LongClickLabel": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "longClickLabel", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Tooltip": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tooltip" in thiz) {
        const val = thiz.tooltip;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Tooltip": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tooltip", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NameFrom": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nameFrom" in thiz) {
        const val = thiz.nameFrom;
        A.store.Enum(
          retPtr,
          [
            "attribute",
            "attributeExplicitlyEmpty",
            "caption",
            "contents",
            "placeholder",
            "popoverAttribute",
            "relatedElement",
            "title",
            "value",
          ].indexOf(val as string)
        );
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NameFrom": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(
        thiz,
        "nameFrom",
        val > 0 && val <= 9
          ? [
              "attribute",
              "attributeExplicitlyEmpty",
              "caption",
              "contents",
              "placeholder",
              "popoverAttribute",
              "relatedElement",
              "title",
              "value",
            ][val - 1]
          : undefined,
        thiz
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationNode_ImageAnnotation": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "imageAnnotation" in thiz) {
        const val = thiz.imageAnnotation;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ImageAnnotation": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "imageAnnotation", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Value": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "value" in thiz) {
        const val = thiz.value;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "get_AutomationNode_HtmlTag": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "htmlTag" in thiz) {
        const val = thiz.htmlTag;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_HtmlTag": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "htmlTag", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_HierarchicalLevel": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "hierarchicalLevel" in thiz) {
        const val = thiz.hierarchicalLevel;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_HierarchicalLevel": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "hierarchicalLevel", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_CaretBounds": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "caretBounds" in thiz) {
        const val = thiz.caretBounds;
        if (typeof val === "undefined") {
          A.store.Bool(retPtr + 20, false);
          A.store.Bool(retPtr + 16, false);
          A.store.Int32(retPtr + 0, 0);
          A.store.Bool(retPtr + 17, false);
          A.store.Int32(retPtr + 4, 0);
          A.store.Bool(retPtr + 18, false);
          A.store.Int32(retPtr + 8, 0);
          A.store.Bool(retPtr + 19, false);
          A.store.Int32(retPtr + 12, 0);
        } else {
          A.store.Bool(retPtr + 20, true);
          A.store.Bool(retPtr + 16, "left" in val ? true : false);
          A.store.Int32(retPtr + 0, val["left"] === undefined ? 0 : (val["left"] as number));
          A.store.Bool(retPtr + 17, "top" in val ? true : false);
          A.store.Int32(retPtr + 4, val["top"] === undefined ? 0 : (val["top"] as number));
          A.store.Bool(retPtr + 18, "width" in val ? true : false);
          A.store.Int32(retPtr + 8, val["width"] === undefined ? 0 : (val["width"] as number));
          A.store.Bool(retPtr + 19, "height" in val ? true : false);
          A.store.Int32(retPtr + 12, val["height"] === undefined ? 0 : (val["height"] as number));
        }
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_CaretBounds": (self: heap.Ref<object>, val: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      const val_ffi = {};

      if (A.load.Bool(val + 16)) {
        val_ffi["left"] = A.load.Int32(val + 0);
      }
      if (A.load.Bool(val + 17)) {
        val_ffi["top"] = A.load.Int32(val + 4);
      }
      if (A.load.Bool(val + 18)) {
        val_ffi["width"] = A.load.Int32(val + 8);
      }
      if (A.load.Bool(val + 19)) {
        val_ffi["height"] = A.load.Int32(val + 12);
      }
      return Reflect.set(thiz, "caretBounds", val_ffi, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_WordStarts": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "wordStarts" in thiz) {
        const val = thiz.wordStarts;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_WordStarts": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "wordStarts", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_WordEnds": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "wordEnds" in thiz) {
        const val = thiz.wordEnds;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_WordEnds": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "wordEnds", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SentenceStarts": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "sentenceStarts" in thiz) {
        const val = thiz.sentenceStarts;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SentenceStarts": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "sentenceStarts", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SentenceEnds": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "sentenceEnds" in thiz) {
        const val = thiz.sentenceEnds;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SentenceEnds": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "sentenceEnds", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NonInlineTextWordStarts": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nonInlineTextWordStarts" in thiz) {
        const val = thiz.nonInlineTextWordStarts;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NonInlineTextWordStarts": (
      self: heap.Ref<object>,
      val: heap.Ref<object>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "nonInlineTextWordStarts", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NonInlineTextWordEnds": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nonInlineTextWordEnds" in thiz) {
        const val = thiz.nonInlineTextWordEnds;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NonInlineTextWordEnds": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "nonInlineTextWordEnds", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Controls": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "controls" in thiz) {
        const val = thiz.controls;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Controls": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "controls", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DescribedBy": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "describedBy" in thiz) {
        const val = thiz.describedBy;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DescribedBy": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "describedBy", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FlowTo": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "flowTo" in thiz) {
        const val = thiz.flowTo;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FlowTo": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "flowTo", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LabelledBy": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "labelledBy" in thiz) {
        const val = thiz.labelledBy;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LabelledBy": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "labelledBy", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ActiveDescendant": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "activeDescendant" in thiz) {
        const val = thiz.activeDescendant;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ActiveDescendant": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "activeDescendant", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ActiveDescendantFor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "activeDescendantFor" in thiz) {
        const val = thiz.activeDescendantFor;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ActiveDescendantFor": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "activeDescendantFor", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_InPageLinkTarget": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "inPageLinkTarget" in thiz) {
        const val = thiz.inPageLinkTarget;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_InPageLinkTarget": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "inPageLinkTarget", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Details": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "details" in thiz) {
        const val = thiz.details;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Details": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "details", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ErrorMessages": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "errorMessages" in thiz) {
        const val = thiz.errorMessages;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ErrorMessages": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "errorMessages", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DetailsFor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "detailsFor" in thiz) {
        const val = thiz.detailsFor;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DetailsFor": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "detailsFor", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ErrorMessageFor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "errorMessageFor" in thiz) {
        const val = thiz.errorMessageFor;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ErrorMessageFor": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "errorMessageFor", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ControlledBy": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "controlledBy" in thiz) {
        const val = thiz.controlledBy;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ControlledBy": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "controlledBy", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DescriptionFor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "descriptionFor" in thiz) {
        const val = thiz.descriptionFor;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DescriptionFor": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "descriptionFor", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FlowFrom": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "flowFrom" in thiz) {
        const val = thiz.flowFrom;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FlowFrom": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "flowFrom", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LabelFor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "labelFor" in thiz) {
        const val = thiz.labelFor;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LabelFor": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "labelFor", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellColumnHeaders": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellColumnHeaders" in thiz) {
        const val = thiz.tableCellColumnHeaders;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellColumnHeaders": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellColumnHeaders", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellRowHeaders": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellRowHeaders" in thiz) {
        const val = thiz.tableCellRowHeaders;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellRowHeaders": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellRowHeaders", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_StandardActions": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "standardActions" in thiz) {
        const val = thiz.standardActions;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_StandardActions": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "standardActions", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_CustomActions": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "customActions" in thiz) {
        const val = thiz.customActions;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_CustomActions": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "customActions", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DefaultActionVerb": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "defaultActionVerb" in thiz) {
        const val = thiz.defaultActionVerb;
        A.store.Enum(
          retPtr,
          ["activate", "check", "click", "clickAncestor", "jump", "open", "press", "select", "uncheck"].indexOf(
            val as string
          )
        );
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DefaultActionVerb": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(
        thiz,
        "defaultActionVerb",
        val > 0 && val <= 9
          ? ["activate", "check", "click", "clickAncestor", "jump", "open", "press", "select", "uncheck"][val - 1]
          : undefined,
        thiz
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationNode_Url": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "url" in thiz) {
        const val = thiz.url;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Url": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "url", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DocUrl": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "docUrl" in thiz) {
        const val = thiz.docUrl;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DocUrl": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "docUrl", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DocTitle": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "docTitle" in thiz) {
        const val = thiz.docTitle;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DocTitle": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "docTitle", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DocLoaded": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "docLoaded" in thiz) {
        const val = thiz.docLoaded;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DocLoaded": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "docLoaded", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DocLoadingProgress": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "docLoadingProgress" in thiz) {
        const val = thiz.docLoadingProgress;
        A.store.Float64(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DocLoadingProgress": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "docLoadingProgress", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ScrollX": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "scrollX" in thiz) {
        const val = thiz.scrollX;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ScrollX": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "scrollX", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ScrollXMin": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "scrollXMin" in thiz) {
        const val = thiz.scrollXMin;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ScrollXMin": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "scrollXMin", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ScrollXMax": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "scrollXMax" in thiz) {
        const val = thiz.scrollXMax;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ScrollXMax": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "scrollXMax", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ScrollY": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "scrollY" in thiz) {
        const val = thiz.scrollY;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ScrollY": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "scrollY", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ScrollYMin": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "scrollYMin" in thiz) {
        const val = thiz.scrollYMin;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ScrollYMin": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "scrollYMin", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ScrollYMax": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "scrollYMax" in thiz) {
        const val = thiz.scrollYMax;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ScrollYMax": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "scrollYMax", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Scrollable": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "scrollable" in thiz) {
        const val = thiz.scrollable;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Scrollable": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "scrollable", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TextSelStart": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "textSelStart" in thiz) {
        const val = thiz.textSelStart;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TextSelStart": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "textSelStart", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TextSelEnd": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "textSelEnd" in thiz) {
        const val = thiz.textSelEnd;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TextSelEnd": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "textSelEnd", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Markers": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "markers" in thiz) {
        const val = thiz.markers;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Markers": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "markers", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_IsSelectionBackward": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "isSelectionBackward" in thiz) {
        const val = thiz.isSelectionBackward;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_IsSelectionBackward": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "isSelectionBackward", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AnchorObject": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "anchorObject" in thiz) {
        const val = thiz.anchorObject;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AnchorObject": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "anchorObject", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AnchorOffset": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "anchorOffset" in thiz) {
        const val = thiz.anchorOffset;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AnchorOffset": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "anchorOffset", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AnchorAffinity": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "anchorAffinity" in thiz) {
        const val = thiz.anchorAffinity;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AnchorAffinity": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "anchorAffinity", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FocusObject": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "focusObject" in thiz) {
        const val = thiz.focusObject;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FocusObject": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "focusObject", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FocusOffset": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "focusOffset" in thiz) {
        const val = thiz.focusOffset;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FocusOffset": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "focusOffset", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FocusAffinity": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "focusAffinity" in thiz) {
        const val = thiz.focusAffinity;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FocusAffinity": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "focusAffinity", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SelectionStartObject": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "selectionStartObject" in thiz) {
        const val = thiz.selectionStartObject;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SelectionStartObject": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "selectionStartObject", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SelectionStartOffset": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "selectionStartOffset" in thiz) {
        const val = thiz.selectionStartOffset;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SelectionStartOffset": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "selectionStartOffset", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SelectionStartAffinity": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "selectionStartAffinity" in thiz) {
        const val = thiz.selectionStartAffinity;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SelectionStartAffinity": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "selectionStartAffinity", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SelectionEndObject": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "selectionEndObject" in thiz) {
        const val = thiz.selectionEndObject;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SelectionEndObject": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "selectionEndObject", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SelectionEndOffset": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "selectionEndOffset" in thiz) {
        const val = thiz.selectionEndOffset;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SelectionEndOffset": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "selectionEndOffset", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SelectionEndAffinity": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "selectionEndAffinity" in thiz) {
        const val = thiz.selectionEndAffinity;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SelectionEndAffinity": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "selectionEndAffinity", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NotUserSelectableStyle": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "notUserSelectableStyle" in thiz) {
        const val = thiz.notUserSelectableStyle;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NotUserSelectableStyle": (
      self: heap.Ref<object>,
      val: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "notUserSelectableStyle", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ValueForRange": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "valueForRange" in thiz) {
        const val = thiz.valueForRange;
        A.store.Float64(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ValueForRange": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "valueForRange", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_MinValueForRange": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "minValueForRange" in thiz) {
        const val = thiz.minValueForRange;
        A.store.Float64(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_MinValueForRange": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "minValueForRange", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_MaxValueForRange": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "maxValueForRange" in thiz) {
        const val = thiz.maxValueForRange;
        A.store.Float64(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_MaxValueForRange": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "maxValueForRange", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_PosInSet": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "posInSet" in thiz) {
        const val = thiz.posInSet;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_PosInSet": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "posInSet", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SetSize": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "setSize" in thiz) {
        const val = thiz.setSize;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SetSize": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "setSize", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableRowCount": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableRowCount" in thiz) {
        const val = thiz.tableRowCount;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableRowCount": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableRowCount", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AriaRowCount": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "ariaRowCount" in thiz) {
        const val = thiz.ariaRowCount;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AriaRowCount": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "ariaRowCount", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableColumnCount": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableColumnCount" in thiz) {
        const val = thiz.tableColumnCount;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableColumnCount": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableColumnCount", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AriaColumnCount": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "ariaColumnCount" in thiz) {
        const val = thiz.ariaColumnCount;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AriaColumnCount": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "ariaColumnCount", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellColumnIndex": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellColumnIndex" in thiz) {
        const val = thiz.tableCellColumnIndex;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellColumnIndex": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellColumnIndex", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellAriaColumnIndex": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellAriaColumnIndex" in thiz) {
        const val = thiz.tableCellAriaColumnIndex;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellAriaColumnIndex": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellAriaColumnIndex", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellColumnSpan": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellColumnSpan" in thiz) {
        const val = thiz.tableCellColumnSpan;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellColumnSpan": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellColumnSpan", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellRowIndex": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellRowIndex" in thiz) {
        const val = thiz.tableCellRowIndex;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellRowIndex": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellRowIndex", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellAriaRowIndex": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellAriaRowIndex" in thiz) {
        const val = thiz.tableCellAriaRowIndex;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellAriaRowIndex": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellAriaRowIndex", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableCellRowSpan": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableCellRowSpan" in thiz) {
        const val = thiz.tableCellRowSpan;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableCellRowSpan": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableCellRowSpan", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableColumnHeader": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableColumnHeader" in thiz) {
        const val = thiz.tableColumnHeader;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableColumnHeader": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableColumnHeader", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableRowHeader": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableRowHeader" in thiz) {
        const val = thiz.tableRowHeader;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableRowHeader": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableRowHeader", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableColumnIndex": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableColumnIndex" in thiz) {
        const val = thiz.tableColumnIndex;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableColumnIndex": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableColumnIndex", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_TableRowIndex": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "tableRowIndex" in thiz) {
        const val = thiz.tableRowIndex;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_TableRowIndex": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "tableRowIndex", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LiveStatus": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "liveStatus" in thiz) {
        const val = thiz.liveStatus;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LiveStatus": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "liveStatus", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LiveRelevant": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "liveRelevant" in thiz) {
        const val = thiz.liveRelevant;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LiveRelevant": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "liveRelevant", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LiveAtomic": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "liveAtomic" in thiz) {
        const val = thiz.liveAtomic;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LiveAtomic": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "liveAtomic", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Busy": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "busy" in thiz) {
        const val = thiz.busy;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Busy": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "busy", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ContainerLiveStatus": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "containerLiveStatus" in thiz) {
        const val = thiz.containerLiveStatus;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ContainerLiveStatus": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "containerLiveStatus", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ContainerLiveRelevant": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "containerLiveRelevant" in thiz) {
        const val = thiz.containerLiveRelevant;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ContainerLiveRelevant": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "containerLiveRelevant", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ContainerLiveAtomic": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "containerLiveAtomic" in thiz) {
        const val = thiz.containerLiveAtomic;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ContainerLiveAtomic": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "containerLiveAtomic", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ContainerLiveBusy": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "containerLiveBusy" in thiz) {
        const val = thiz.containerLiveBusy;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ContainerLiveBusy": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "containerLiveBusy", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_IsButton": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "isButton" in thiz) {
        const val = thiz.isButton;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_IsButton": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "isButton", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_IsCheckBox": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "isCheckBox" in thiz) {
        const val = thiz.isCheckBox;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_IsCheckBox": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "isCheckBox", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_IsComboBox": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "isComboBox" in thiz) {
        const val = thiz.isComboBox;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_IsComboBox": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "isComboBox", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_IsImage": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "isImage" in thiz) {
        const val = thiz.isImage;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_IsImage": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "isImage", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_HasHiddenOffscreenNodes": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "hasHiddenOffscreenNodes" in thiz) {
        const val = thiz.hasHiddenOffscreenNodes;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_HasHiddenOffscreenNodes": (
      self: heap.Ref<object>,
      val: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "hasHiddenOffscreenNodes", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AutoComplete": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "autoComplete" in thiz) {
        const val = thiz.autoComplete;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AutoComplete": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "autoComplete", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ClassName": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "className" in thiz) {
        const val = thiz.className;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ClassName": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "className", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Modal": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "modal" in thiz) {
        const val = thiz.modal;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Modal": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "modal", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_HtmlAttributes": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "htmlAttributes" in thiz) {
        const val = thiz.htmlAttributes;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_HtmlAttributes": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "htmlAttributes", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_InputType": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "inputType" in thiz) {
        const val = thiz.inputType;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_InputType": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "inputType", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AccessKey": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "accessKey" in thiz) {
        const val = thiz.accessKey;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AccessKey": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "accessKey", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AriaInvalidValue": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "ariaInvalidValue" in thiz) {
        const val = thiz.ariaInvalidValue;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AriaInvalidValue": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "ariaInvalidValue", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Display": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "display" in thiz) {
        const val = thiz.display;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Display": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "display", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ImageDataUrl": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "imageDataUrl" in thiz) {
        const val = thiz.imageDataUrl;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ImageDataUrl": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "imageDataUrl", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Language": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "language" in thiz) {
        const val = thiz.language;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Language": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "language", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_DetectedLanguage": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "detectedLanguage" in thiz) {
        const val = thiz.detectedLanguage;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_DetectedLanguage": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "detectedLanguage", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_HasPopup": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "hasPopup" in thiz) {
        const val = thiz.hasPopup;
        A.store.Enum(retPtr, ["false", "true", "menu", "listbox", "tree", "grid", "dialog"].indexOf(val as string));
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_HasPopup": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(
        thiz,
        "hasPopup",
        val > 0 && val <= 7 ? ["false", "true", "menu", "listbox", "tree", "grid", "dialog"][val - 1] : undefined,
        thiz
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationNode_Restriction": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "restriction" in thiz) {
        const val = thiz.restriction;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Restriction": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "restriction", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Checked": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "checked" in thiz) {
        const val = thiz.checked;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Checked": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "checked", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_InnerHtml": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "innerHtml" in thiz) {
        const val = thiz.innerHtml;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_InnerHtml": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "innerHtml", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Color": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "color" in thiz) {
        const val = thiz.color;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Color": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "color", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_BackgroundColor": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "backgroundColor" in thiz) {
        const val = thiz.backgroundColor;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_BackgroundColor": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "backgroundColor", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_ColorValue": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "colorValue" in thiz) {
        const val = thiz.colorValue;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_ColorValue": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "colorValue", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Subscript": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "subscript" in thiz) {
        const val = thiz.subscript;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Subscript": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "subscript", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Superscript": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "superscript" in thiz) {
        const val = thiz.superscript;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Superscript": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "superscript", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Bold": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "bold" in thiz) {
        const val = thiz.bold;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Bold": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "bold", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Italic": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "italic" in thiz) {
        const val = thiz.italic;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Italic": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "italic", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Underline": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "underline" in thiz) {
        const val = thiz.underline;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Underline": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "underline", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LineThrough": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "lineThrough" in thiz) {
        const val = thiz.lineThrough;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LineThrough": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "lineThrough", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Selected": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "selected" in thiz) {
        const val = thiz.selected;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Selected": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "selected", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FontSize": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "fontSize" in thiz) {
        const val = thiz.fontSize;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FontSize": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "fontSize", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FontFamily": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "fontFamily" in thiz) {
        const val = thiz.fontFamily;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FontFamily": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "fontFamily", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NonAtomicTextFieldRoot": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nonAtomicTextFieldRoot" in thiz) {
        const val = thiz.nonAtomicTextFieldRoot;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NonAtomicTextFieldRoot": (
      self: heap.Ref<object>,
      val: heap.Ref<boolean>
    ): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "nonAtomicTextFieldRoot", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_AriaCurrentState": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "ariaCurrentState" in thiz) {
        const val = thiz.ariaCurrentState;
        A.store.Enum(retPtr, ["false", "true", "page", "step", "location", "date", "time"].indexOf(val as string));
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AriaCurrentState": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(
        thiz,
        "ariaCurrentState",
        val > 0 && val <= 7 ? ["false", "true", "page", "step", "location", "date", "time"][val - 1] : undefined,
        thiz
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationNode_InvalidState": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "invalidState" in thiz) {
        const val = thiz.invalidState;
        A.store.Enum(retPtr, ["false", "true"].indexOf(val as string));
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_InvalidState": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "invalidState", val > 0 && val <= 2 ? ["false", "true"][val - 1] : undefined, thiz)
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationNode_AppId": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "appId" in thiz) {
        const val = thiz.appId;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_AppId": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "appId", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Children": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "children" in thiz) {
        const val = thiz.children;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Children": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "children", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_Parent": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "parent" in thiz) {
        const val = thiz.parent;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Parent": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "parent", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_FirstChild": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "firstChild" in thiz) {
        const val = thiz.firstChild;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_FirstChild": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "firstChild", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_LastChild": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "lastChild" in thiz) {
        const val = thiz.lastChild;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_LastChild": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "lastChild", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_PreviousSibling": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "previousSibling" in thiz) {
        const val = thiz.previousSibling;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_PreviousSibling": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "previousSibling", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NextSibling": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nextSibling" in thiz) {
        const val = thiz.nextSibling;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NextSibling": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "nextSibling", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_PreviousOnLine": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "previousOnLine" in thiz) {
        const val = thiz.previousOnLine;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_PreviousOnLine": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "previousOnLine", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NextOnLine": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nextOnLine" in thiz) {
        const val = thiz.nextOnLine;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NextOnLine": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "nextOnLine", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_PreviousFocus": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "previousFocus" in thiz) {
        const val = thiz.previousFocus;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_PreviousFocus": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "previousFocus", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NextFocus": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nextFocus" in thiz) {
        const val = thiz.nextFocus;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NextFocus": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "nextFocus", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_PreviousWindowFocus": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "previousWindowFocus" in thiz) {
        const val = thiz.previousWindowFocus;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_PreviousWindowFocus": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "previousWindowFocus", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_NextWindowFocus": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "nextWindowFocus" in thiz) {
        const val = thiz.nextWindowFocus;
        A.store.Ref(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_NextWindowFocus": (self: heap.Ref<object>, val: heap.Ref<object>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "nextWindowFocus", A.H.get<object>(val), thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_IndexInParent": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "indexInParent" in thiz) {
        const val = thiz.indexInParent;
        A.store.Int32(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_IndexInParent": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "indexInParent", val, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "get_AutomationNode_SortDirection": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "sortDirection" in thiz) {
        const val = thiz.sortDirection;
        A.store.Enum(retPtr, ["unsorted", "ascending", "descending", "other"].indexOf(val as string));
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_SortDirection": (self: heap.Ref<object>, val: number): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(
        thiz,
        "sortDirection",
        val > 0 && val <= 4 ? ["unsorted", "ascending", "descending", "other"][val - 1] : undefined,
        thiz
      )
        ? A.H.TRUE
        : A.H.FALSE;
    },
    "get_AutomationNode_Clickable": (self: heap.Ref<any>, retPtr: Pointer): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);
      if (thiz && "clickable" in thiz) {
        const val = thiz.clickable;
        A.store.Bool(retPtr, val);
        return A.H.TRUE;
      }

      return A.H.FALSE;
    },
    "set_AutomationNode_Clickable": (self: heap.Ref<object>, val: heap.Ref<boolean>): heap.Ref<boolean> => {
      const thiz = A.H.get<object>(self);

      return Reflect.set(thiz, "clickable", val === A.H.TRUE, thiz) ? A.H.TRUE : A.H.FALSE;
    },
    "constof_DescriptionFromType": (ref: heap.Ref<string>): number => {
      const idx = [
        "ariaDescription",
        "attributeExplicitlyEmpty",
        "buttonLabel",
        "popoverAttribute",
        "relatedElement",
        "rubyAnnotation",
        "summary",
        "svgDescElement",
        "tableCaption",
        "title",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_IntentInputEventType": (ref: heap.Ref<string>): number => {
      const idx = [
        "insertText",
        "insertLineBreak",
        "insertParagraph",
        "insertOrderedList",
        "insertUnorderedList",
        "insertHorizontalRule",
        "insertFromPaste",
        "insertFromDrop",
        "insertFromYank",
        "insertTranspose",
        "insertReplacementText",
        "insertCompositionText",
        "deleteWordBackward",
        "deleteWordForward",
        "deleteSoftLineBackward",
        "deleteSoftLineForward",
        "deleteHardLineBackward",
        "deleteHardLineForward",
        "deleteContentBackward",
        "deleteContentForward",
        "deleteByCut",
        "deleteByDrag",
        "historyUndo",
        "historyRedo",
        "formatBold",
        "formatItalic",
        "formatUnderline",
        "formatStrikeThrough",
        "formatSuperscript",
        "formatSubscript",
        "formatJustifyCenter",
        "formatJustifyFull",
        "formatJustifyRight",
        "formatJustifyLeft",
        "formatIndent",
        "formatOutdent",
        "formatRemove",
        "formatSetBlockTextDirection",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_MarkerType": (ref: heap.Ref<string>): number => {
      const idx = ["spelling", "grammar", "textMatch", "activeSuggestion", "suggestion", "highlight"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_Restriction": (ref: heap.Ref<string>): number => {
      const idx = ["disabled", "readOnly"].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_SetDocumentSelectionParams": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 18, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Bool(ptr + 16, false);
        A.store.Int32(ptr + 4, 0);
        A.store.Ref(ptr + 8, undefined);
        A.store.Bool(ptr + 17, false);
        A.store.Int32(ptr + 12, 0);
      } else {
        A.store.Bool(ptr + 18, true);
        A.store.Ref(ptr + 0, x["anchorObject"]);
        A.store.Bool(ptr + 16, "anchorOffset" in x ? true : false);
        A.store.Int32(ptr + 4, x["anchorOffset"] === undefined ? 0 : (x["anchorOffset"] as number));
        A.store.Ref(ptr + 8, x["focusObject"]);
        A.store.Bool(ptr + 17, "focusOffset" in x ? true : false);
        A.store.Int32(ptr + 12, x["focusOffset"] === undefined ? 0 : (x["focusOffset"] as number));
      }
    },
    "load_SetDocumentSelectionParams": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["anchorObject"] = A.load.Ref(ptr + 0, undefined);
      if (A.load.Bool(ptr + 16)) {
        x["anchorOffset"] = A.load.Int32(ptr + 4);
      } else {
        delete x["anchorOffset"];
      }
      x["focusObject"] = A.load.Ref(ptr + 8, undefined);
      if (A.load.Bool(ptr + 17)) {
        x["focusOffset"] = A.load.Int32(ptr + 12);
      } else {
        delete x["focusOffset"];
      }
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_StateType": (ref: heap.Ref<string>): number => {
      const idx = [
        "autofillAvailable",
        "collapsed",
        "default",
        "editable",
        "expanded",
        "focusable",
        "focused",
        "horizontal",
        "hovered",
        "ignored",
        "invisible",
        "linked",
        "multiline",
        "multiselectable",
        "offscreen",
        "protected",
        "required",
        "richlyEditable",
        "vertical",
        "visited",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },
    "constof_TreeChangeType": (ref: heap.Ref<string>): number => {
      const idx = [
        "nodeCreated",
        "subtreeCreated",
        "nodeChanged",
        "textChanged",
        "nodeRemoved",
        "subtreeUpdateEnd",
      ].indexOf(A.H.get(ref));
      return idx < 0 ? 0 : idx + 1;
    },

    "store_TreeChange": (ptr: Pointer, ref: heap.Ref<any>) => {
      const x = A.H.get<any>(ref);

      if (typeof x === "undefined") {
        A.store.Bool(ptr + 8, false);
        A.store.Ref(ptr + 0, undefined);
        A.store.Enum(ptr + 4, -1);
      } else {
        A.store.Bool(ptr + 8, true);
        A.store.Ref(ptr + 0, x["target"]);
        A.store.Enum(
          ptr + 4,
          ["nodeCreated", "subtreeCreated", "nodeChanged", "textChanged", "nodeRemoved", "subtreeUpdateEnd"].indexOf(
            x["type"] as string
          )
        );
      }
    },
    "load_TreeChange": (ptr: Pointer, create: heap.Ref<boolean>, ref: heap.Ref<any>): heap.Ref<any> => {
      const x = create === A.H.TRUE ? {} : A.H.get<any>(ref);

      x["target"] = A.load.Ref(ptr + 0, undefined);
      x["type"] = A.load.Enum(ptr + 4, [
        "nodeCreated",
        "subtreeCreated",
        "nodeChanged",
        "textChanged",
        "nodeRemoved",
        "subtreeUpdateEnd",
      ]);
      return create === A.H.TRUE ? A.H.push(x) : ref;
    },
    "constof_TreeChangeObserverFilter": (ref: heap.Ref<string>): number => {
      const idx = ["noTreeChanges", "liveRegionTreeChanges", "textMarkerChanges", "allTreeChanges"].indexOf(
        A.H.get(ref)
      );
      return idx < 0 ? 0 : idx + 1;
    },
    "has_AddTreeChangeObserver": (): heap.Ref<boolean> => {
      if (WEBEXT?.automation && "addTreeChangeObserver" in WEBEXT?.automation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_AddTreeChangeObserver": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automation.addTreeChangeObserver);
    },
    "call_AddTreeChangeObserver": (retPtr: Pointer, filter: number, observer: heap.Ref<object>): void => {
      const _ret = WEBEXT.automation.addTreeChangeObserver(
        filter > 0 && filter <= 4
          ? ["noTreeChanges", "liveRegionTreeChanges", "textMarkerChanges", "allTreeChanges"][filter - 1]
          : undefined,
        A.H.get<object>(observer)
      );
    },
    "try_AddTreeChangeObserver": (
      retPtr: Pointer,
      errPtr: Pointer,
      filter: number,
      observer: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automation.addTreeChangeObserver(
          filter > 0 && filter <= 4
            ? ["noTreeChanges", "liveRegionTreeChanges", "textMarkerChanges", "allTreeChanges"][filter - 1]
            : undefined,
          A.H.get<object>(observer)
        );
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetAccessibilityFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.automation && "getAccessibilityFocus" in WEBEXT?.automation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetAccessibilityFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automation.getAccessibilityFocus);
    },
    "call_GetAccessibilityFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automation.getAccessibilityFocus(A.H.get<object>(callback));
    },
    "try_GetAccessibilityFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automation.getAccessibilityFocus(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetDesktop": (): heap.Ref<boolean> => {
      if (WEBEXT?.automation && "getDesktop" in WEBEXT?.automation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetDesktop": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automation.getDesktop);
    },
    "call_GetDesktop": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automation.getDesktop(A.H.get<object>(callback));
    },
    "try_GetDesktop": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automation.getDesktop(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetFocus": (): heap.Ref<boolean> => {
      if (WEBEXT?.automation && "getFocus" in WEBEXT?.automation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetFocus": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automation.getFocus);
    },
    "call_GetFocus": (retPtr: Pointer, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automation.getFocus(A.H.get<object>(callback));
    },
    "try_GetFocus": (retPtr: Pointer, errPtr: Pointer, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automation.getFocus(A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_GetTree": (): heap.Ref<boolean> => {
      if (WEBEXT?.automation && "getTree" in WEBEXT?.automation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_GetTree": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automation.getTree);
    },
    "call_GetTree": (retPtr: Pointer, tabId: number, callback: heap.Ref<object>): void => {
      const _ret = WEBEXT.automation.getTree(tabId, A.H.get<object>(callback));
    },
    "try_GetTree": (retPtr: Pointer, errPtr: Pointer, tabId: number, callback: heap.Ref<object>): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automation.getTree(tabId, A.H.get<object>(callback));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_RemoveTreeChangeObserver": (): heap.Ref<boolean> => {
      if (WEBEXT?.automation && "removeTreeChangeObserver" in WEBEXT?.automation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_RemoveTreeChangeObserver": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automation.removeTreeChangeObserver);
    },
    "call_RemoveTreeChangeObserver": (retPtr: Pointer, observer: heap.Ref<object>): void => {
      const _ret = WEBEXT.automation.removeTreeChangeObserver(A.H.get<object>(observer));
    },
    "try_RemoveTreeChangeObserver": (
      retPtr: Pointer,
      errPtr: Pointer,
      observer: heap.Ref<object>
    ): heap.Ref<boolean> => {
      try {
        const _ret = WEBEXT.automation.removeTreeChangeObserver(A.H.get<object>(observer));
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
    "has_SetDocumentSelection": (): heap.Ref<boolean> => {
      if (WEBEXT?.automation && "setDocumentSelection" in WEBEXT?.automation) {
        return A.H.TRUE;
      }
      return A.H.FALSE;
    },
    "func_SetDocumentSelection": (fn: Pointer): void => {
      A.store.Ref(fn, WEBEXT.automation.setDocumentSelection);
    },
    "call_SetDocumentSelection": (retPtr: Pointer, params: Pointer): void => {
      const params_ffi = {};

      params_ffi["anchorObject"] = A.load.Ref(params + 0, undefined);
      if (A.load.Bool(params + 16)) {
        params_ffi["anchorOffset"] = A.load.Int32(params + 4);
      }
      params_ffi["focusObject"] = A.load.Ref(params + 8, undefined);
      if (A.load.Bool(params + 17)) {
        params_ffi["focusOffset"] = A.load.Int32(params + 12);
      }

      const _ret = WEBEXT.automation.setDocumentSelection(params_ffi);
    },
    "try_SetDocumentSelection": (retPtr: Pointer, errPtr: Pointer, params: Pointer): heap.Ref<boolean> => {
      try {
        const params_ffi = {};

        params_ffi["anchorObject"] = A.load.Ref(params + 0, undefined);
        if (A.load.Bool(params + 16)) {
          params_ffi["anchorOffset"] = A.load.Int32(params + 4);
        }
        params_ffi["focusObject"] = A.load.Ref(params + 8, undefined);
        if (A.load.Bool(params + 17)) {
          params_ffi["focusOffset"] = A.load.Int32(params + 12);
        }

        const _ret = WEBEXT.automation.setDocumentSelection(params_ffi);
        return A.H.TRUE;
      } catch (err: any) {
        A.store.Ref(errPtr, err);
        return A.H.FALSE;
      }
    },
  };
});
